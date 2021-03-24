package wsApp

import (
	"encoding/json"
	"log"
	"os"
	"ws-rest-test/pkg/user"
)

type MediaLink struct {
	Id         int64  `json:"id"`
	MediaUrl   string `json:"mediaUrl"`
	PreviewUrl string `json:"previewUrl"`
}

func (w *wsApp) CreateMediaLink(user *user.User) string {
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

	return url.MediaUrl
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
	//fmt.Println(url)
	w.rest_client.DoPut(url, data, headers)
	//fmt.Println(string(out))

	//url := MediaLink{}
	//err := json.Unmarshal(out, &url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(url.MediaUrl)
}
