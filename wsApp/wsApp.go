package wsApp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync/atomic"
	"time"

	"ws-rest-test/pkg/rest_client"
	"ws-rest-test/pkg/user"

	guuid "github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var jobCounter uint64

type wsApp struct {
	Config      *Config
	rest_client *rest_client.Rest_client
	ops         uint64
}

func NewWsApp(config *Config) *wsApp {
	return &wsApp{
		Config: config,
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

func (w *wsApp) ReportCounter(ctx context.Context) {
	current := jobCounter
	for {
		select {
		case <-ctx.Done():
			return
		default:
			{
				logrus.Println(jobCounter-current, " messages sent in ", w.Config.Report_timeout, "(ms)")
				current = jobCounter
				time.Sleep(time.Duration(w.Config.Report_timeout) * time.Millisecond)
			}
		}
	}
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
}

func (w *wsApp) SetOnlineStatus(user *user.User) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoGet(w.rest_client.Config.Url+"/api/v1/chats/online", headers)
}

func (w *wsApp) GetWSPrivate(user *user.User) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoGet(w.rest_client.Config.Url+"/api/v1/tokens/ws/private", headers)
}

func (w *wsApp) GetEventState(user *user.User) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoGet(w.rest_client.Config.Url+"/api/v1/events/state", headers)
}

func (w *wsApp) PutInfo(user *user.User, name string, age string) {
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
}

func (w *wsApp) CreateChat(user *user.User, phone string) int64 {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
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

	return chatID.Id
}

func (w *wsApp) InstantStatusTyping(user *user.User, chatID string) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoPut(w.rest_client.Config.Url+"/api/v1/chats/"+chatID+"/instantStatus/Typing", nil, headers)
}

func (w *wsApp) SendMessage(sender *user.User, recipient *user.User, chatID string) {

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
	w.rest_client.DoPost(w.rest_client.Config.Url+"/api/v1/chats/"+chatID+"/messages", []byte(message), headers)

}

func (w *wsApp) RunUser(users *pairUsers, ctx context.Context) {
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
			if w.Config.Run_message_test == 1 {
				time.Sleep(time.Duration(w.Config.Message_test_timeout) * time.Millisecond)
				w.SendMessage(users.user1, users.user2, chatID)
			}
			if w.Config.Run_video_test == 1 {
				time.Sleep(time.Duration(w.Config.Video_test_timeout) * time.Millisecond)
				url := w.CreateMediaLink(users.user1)
				w.UploadVideo(url.MediaUrl, w.Config.Path_video_file)
				w.UploadImage(url.PreviewUrl, w.Config.Path_image_preview_file)
				if w.Config.Send_video_message_to_chat == 1 {
					w.MultySendVideo(users.user1, chatID, &url)
				}
			}
			atomic.AddUint64(&jobCounter, 1)
		}
	}

}

type pairUsers struct {
	user1 *user.User
	user2 *user.User
}

func (w *wsApp) Run() {
	users := make([]*pairUsers, 0, w.Config.Sessions)
	w.rest_client = rest_client.NewRestClient(w.Config.Rest_client)
	ctx, finish := context.WithCancel(context.Background())
	for i := 0; i < w.Config.Sessions; i++ {

		if i%100 == 0 {
			log.Printf("...%d", i)
		}

		user1 := user.NewUser()
		time.Sleep(time.Duration(w.Config.Create_user_timeout) * time.Millisecond)
		w.CreateToken(user1)
		user2 := user.NewUser()
		time.Sleep(time.Duration(w.Config.Create_user_timeout) * time.Millisecond)
		w.CreateToken(user2)
		users = append(users, &pairUsers{user1, user2})

	}
	fmt.Println("Start test:")
	for _, users := range users {
		go w.RunUser(users, ctx)
	}
	go w.ReportCounter(ctx)

	fmt.Scanln()
	finish()
	time.Sleep(time.Duration(2000) * time.Millisecond)

}
