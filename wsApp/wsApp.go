package wsApp

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"ws-rest-test/pkg/login"
	"ws-rest-test/pkg/rest_client"
	"ws-rest-test/pkg/user"
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

func (w *wsApp) createToken(username string) string {

	login := login.NewLogin(username)
	token, err := login.GetToken(w.config.secret)
	if err != nil {
		log.Fatalf("Error generate token process")
	}
	return token
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

func (w *wsApp) CreateToken(user *user.User) {
	body := `{
		"phone": "` + user.GetTel() + `",
		"deviceId": "` + user.GetUUID() + `"
	}`
	w.rest_client.DoPost("/api/v1/pub/signup/code/send", []byte(body))

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

	bodyConfirm := w.rest_client.DoPost("/api/v1/pub/signup/code/confirm", []byte(body))

	confirmJSONobj := confirmJSONw{}
	jsonErr := json.Unmarshal(bodyConfirm, &confirmJSONobj)
	if jsonErr != nil {
		log.Println("CreateToken:unmarshal: ")
		log.Fatal(jsonErr)
	}
	//fmt.Println(confirmJSONobj.Auth.Access.Token)
	user.SetToken(confirmJSONobj.Auth.Access.Token)
}

func (w *wsApp) SetName(user *user.User, name string) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	body := `{
		"name": "` + name + `"
	  }`
	w.rest_client.DoPatch("/api/v1/profile/name", []byte(body), headers)
}

func (w *wsApp) GetResources(user *user.User) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoGet("/api/v1/pub/gp/resources", headers)
	//fmt.Println(string(resp))
}

func (w *wsApp) SetOnlineStatus(user *user.User) {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	w.rest_client.DoGet("/api/v1/chats/online", headers)
	//fmt.Println(string(resp))
}

func (w *wsApp) RunUser(user *user.User) {

	//fmt.Println(user.GetToken())
	w.SetName(user, "TEST")
	w.GetResources(user)
	w.SetOnlineStatus(user)

}

func (w *wsApp) Run() {
	users := make([]*user.User, 0, w.config.sessions)
	config := rest_client.NewConfig()
	w.rest_client = rest_client.NewRestClient(config)
	for i := 0; i < w.config.sessions; i++ {

		if i%100 == 0 {
			fmt.Printf("...%d", i)
		}

		time.Sleep(time.Duration(w.config.create_user_timeout) * time.Millisecond)
		user := user.NewUser()
		w.CreateToken(user)
		users = append(users, user)
	}
	fmt.Println("Start test:")
	for _, userA := range users {
		go w.RunUser(userA)
	}

	//for i := 0; i < w.config.sessions; i++ {
	//}
	//for i := 1; i < 10; i++ {
	time.Sleep(time.Duration(w.config.wait_end_work_timeout) * time.Microsecond)
	//fmt.Println("ops:", atomic.LoadUint64(&w.ops))
	//}
	fmt.Println("Finished")
	fmt.Scanln()

}
