package wsApp

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"ws-rest-test/pkg/user"

	guuid "github.com/google/uuid"
)

type MediaLink struct {
	Id         int64  `json:"id"`
	MediaUrl   string `json:"mediaUrl"`
	PreviewUrl string `json:"previewUrl"`
}

func (w *wsApp) CreateMediaLink(user *user.User) MediaLink {
	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
	}
	//fmt.Println(phone)
	out := w.rest_client.DoPost(w.rest_client.Config.Url+"/api/v1/medias/create/?mediaType=Video", nil, headers)

	url := MediaLink{}
	err := json.Unmarshal(out, &url)
	if err != nil {
		log.Println("Unmarshal:MediaLink")
		log.Fatal(err)
	}
	//fmt.Println(url.MediaUrl)

	return url
}

func (w *wsApp) UploadVideo(url string, file string) {
	data, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	headers := map[string]string{
		"Content-Type": "video/mp4",
	}
	w.rest_client.DoPut(url, data, headers)
}

func (w *wsApp) UploadImage(url string, file string) {
	data, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	headers := map[string]string{
		"Content-Type": "image/jpeg",
	}
	w.rest_client.DoPut(url, data, headers)
}

// Send video link into shared chat
func (w *wsApp) MultySendVideo(user *user.User, chatID string, link *MediaLink) {

	body := `{
		"extId": "` + guuid.New().String() + `",
		"item": "0",
		"mediaId": "` + strconv.FormatInt(link.Id, 10) + `",
		"lifeCycle": {
		  "secondsAfterPlay": 86400
		},
		"locale": "RU",
		"recorded": "2021-03-22T13:50:23.333Z",
		"text": "Video Message By User1, 0 part"
	  }`

	headers := map[string]string{
		"Authorization": "Bearer " + user.GetToken(),
		"Content-Type":  "application/json",
	}
	//fmt.Println(message)
	w.rest_client.DoPost(w.rest_client.Config.Url+"/api/v1/chats/"+chatID+"/messages", []byte(body), headers)
	//fmt.Println("Received: " + string(out) + " \n ChatID: " + chatID)

}
