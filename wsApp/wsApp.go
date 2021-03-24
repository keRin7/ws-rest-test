package wsApp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"ws-rest-test/pkg/rest_client"
	"ws-rest-test/pkg/user"

	guuid "github.com/google/uuid"
)

type wsApp struct {
	config      *config
	rest_client *rest_client.Rest_client
	ops         uint64
}

func NewWsApp(config *config) *wsApp {
	return &wsApp{
		config: config,
		ops:    0,
	}
}

type confirmJSONAuth struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

type confirmJSON struct {
	Access  confirmJSONAuth `json:"access"`
	Refresh confirmJSONAuth `json:"refresh"`
}

type confirmJSONw struct {
	Auth confirmJSON `json:"auth"`
}

type chatID struct {
	Id int64 `json:"id"`
}

func (w *wsApp) CreateToken(user *user.User) {

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	body := `{
		"phone": "` + user.GetTel() + `",
		"deviceId": "` + user.GetUUID() + `"
	}`
	w.rest_client.DoPost(w.rest_client.Config.Url+"/api/v1/pub/signup/code/send", []byte(body), headers)
	body = `{
			"phone": "` + user.GetTel() + `",
			"deviceId": "` + user.GetUUID() + `",
			"code": 1111,
			"notificationToken": null,
			"deviceInfo": {
			  "appVersion": "1.0.172",
			  "model": "iPhone X",
			  "osVersion": 12.1,
			  "type": "IOS",
			  "vendor": "Apple Inc."
			}
		  }`
	confirmJSONobj := confirmJSONw{}
	for {
		bodyConfirm := w.rest_client.DoPost(w.rest_client.Config.Url+"/api/v1/pub/signup/code/confirm", []byte(body), headers)

		//fmt.Println(string(bodyConfirm))

		if bodyConfirm != nil {
			jsonErr := json.Unmarshal(bodyConfirm, &confirmJSONobj)
			if jsonErr != nil {
				log.Println("CreateToken:unmarshal: ")
				log.Fatal(jsonErr)
			}
			break
		} else {
			time.Sleep(time.Duration(1000) * time.Millisecond)
		}
	}
	//fmt.Println(confirmJSONobj.Auth.Access.Token)
	user.SetToken(confirmJSONobj.Auth.Access.Token)
}

func (w *wsApp) SetName(user *user.User, name string) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
		"Content-Type":  "application/json",
	}
	body := `{
		"name": "` + name + `"
	  }`
	w.rest_client.DoPatch(w.rest_client.Config.Url+"/api/v1/profile/name", []byte(body), headers)
}

func (w *wsApp) GetResources(user *user.User) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoGet(w.rest_client.Config.Url+"/api/v1/pub/gp/resources", headers)
	//fmt.Println(string(resp))
}

func (w *wsApp) SetOnlineStatus(user *user.User) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoGet(w.rest_client.Config.Url+"/api/v1/chats/online", headers)
	//fmt.Println(string(resp))
}

func (w *wsApp) GetWSPrivate(user *user.User) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoGet(w.rest_client.Config.Url+"/api/v1/tokens/ws/private", headers)
	//fmt.Println(string(resp))

}

func (w *wsApp) GetEventState(user *user.User) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoGet(w.rest_client.Config.Url+"/api/v1/events/state", headers)
	//fmt.Println(string(resp))
}

func (w *wsApp) PutInfo(user *user.User, name string, age string) {
	//fmt.Println(user.GetToken())
	body := `{
		"age": ` + age + `,
		"forceUpdate": true,
		"gender": "Female",
		"locale": "ru-ru",
		"name": "` + name + `",
		"previousVersion": "string",
		"updateChatLocales": true
	  }`
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
		"Content-Type":  "application/json",
	}

	w.rest_client.DoPut(w.rest_client.Config.Url+"/api/v1/profile/props", bytes.NewReader([]byte(body)), headers)
	//fmt.Println(string(out))
}

func (w *wsApp) CreateChat(user *user.User, phone string) int64 {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	//fmt.Println(phone)
	chatID := chatID{}
	for {
		out := w.rest_client.DoPost(w.rest_client.Config.Url+"/api/v1/chats/by/phone/"+phone, nil, headers)

		if out != nil {
			err := json.Unmarshal(out, &chatID)
			if err != nil {
				log.Fatal(err)
			}
			break
		} else {
			time.Sleep(time.Duration(1000) * time.Millisecond)
		}
	}

	//fmt.Println(chatID.Id)
	return chatID.Id
}

func (w *wsApp) InstantStatusTyping(user *user.User, chatID string) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoPut(w.rest_client.Config.Url+"/api/v1/chats/"+chatID+"/instantStatus/Typing", nil, headers)
}

func (w *wsApp) SendMessage(sender *user.User, recipient *user.User, chatID string) {
	//var out []byte
	//defer log.Println(string(out))

	message := `{
		"extId": "` + guuid.New().String() + `",
		"item": 0,
		"lifeCycle": {
		  "secondsAfterPlay": 86400
		},
		"locale": "RU",
		"recorded": "2021-03-22T13:50:23.333Z",
		"text": "Text Message 0 By User-1"
	  }`

	headers := map[string]string{
		"Authorization": "Bearer " + sender.GetToken(),
		"Content-Type":  "application/json",
	}
	//fmt.Println(message)
	w.rest_client.DoPost(w.rest_client.Config.Url+"/api/v1/chats/"+chatID+"/messages", []byte(message), headers)
	//fmt.Println("Received: " + string(out) + " \n ChatID: " + chatID)

}

func (w *wsApp) RunUser(users *pairUsers, ctx context.Context) {
	//defer fmt.Println("Finish")
	//fmt.Println(user.GetToken())
	w.SetName(users.user1, "TEST123")
	w.SetName(users.user2, "TEST456")
	w.GetResources(users.user1)
	w.GetResources(users.user2)
	w.SetOnlineStatus(users.user1)
	w.SetOnlineStatus(users.user2)
	w.GetWSPrivate(users.user1)
	w.GetWSPrivate(users.user2)
	w.PutInfo(users.user2, "Kasandra", "35")
	w.GetEventState(users.user1)
	chatID := strconv.FormatInt(w.CreateChat(users.user1, users.user2.GetTel()), 10)
	w.InstantStatusTyping(users.user1, chatID)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if w.config.run_message_test == 1 {
				time.Sleep(time.Duration(w.config.message_test_timeout) * time.Millisecond)
				w.SendMessage(users.user1, users.user2, chatID)
			}
			if w.config.run_video_test == 1 {
				time.Sleep(time.Duration(w.config.video_test_timeout) * time.Millisecond)
				url := w.CreateMediaLink(users.user1)
				w.UploadVideo(url, w.config.path_video_file)
			}
		}
	}

}

type pairUsers struct {
	user1 *user.User
	user2 *user.User
}

func (w *wsApp) Run() {
	users := make([]*pairUsers, 0, w.config.sessions)
	config := rest_client.NewConfig()
	w.rest_client = rest_client.NewRestClient(config)
	ctx, finish := context.WithCancel(context.Background())
	for i := 0; i < w.config.sessions; i++ {

		if i%100 == 0 {
			log.Printf("...%d", i)
		}

		user1 := user.NewUser()
		time.Sleep(time.Duration(w.config.create_user_timeout) * time.Millisecond)
		w.CreateToken(user1)
		user2 := user.NewUser()
		time.Sleep(time.Duration(w.config.create_user_timeout) * time.Millisecond)
		w.CreateToken(user2)
		users = append(users, &pairUsers{user1, user2})

	}
	fmt.Println("Start test:")
	for _, users := range users {
		go w.RunUser(users, ctx)
	}

	//go func(ctx context.Context) {
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			return
	//		default:
	//			time.Sleep(time.Duration(2000) * time.Millisecond)
	//			log.Println(runtime.NumGoroutine())
	//		}
	//	}
	//}(ctx)

	fmt.Scanln()
	finish()
	time.Sleep(time.Duration(2000) * time.Millisecond)

}
