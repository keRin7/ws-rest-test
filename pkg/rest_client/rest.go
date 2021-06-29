package rest_client

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Rest_client struct {
	Config     *Config
	httpClient *http.Client
}

func NewRestClient(config *Config) *Rest_client {
	return &Rest_client{
		Config:     config,
		httpClient: &http.Client{},
	}
}

func (r *Rest_client) DoGet(path string, headers map[string]string) []byte {

	req, err := http.NewRequest(http.MethodGet, path, nil)

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := r.httpClient.Do(req)
	if err != nil {
		logrus.Printf("DoGet: %s", path)
		logrus.Fatal(err)
	}
	defer resp.Body.Close()

	if (resp.Status != "200 OK") && (resp.Status != "201 Created") {
		logrus.Println("GET:", resp.Status, " path: ", path)
	} else {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Printf("DoGet:ReadAll: %s", path)
			logrus.Fatal(err)
		}
		return body
	}
	return nil
}

func (r *Rest_client) DoPost(path string, body []byte, headers map[string]string) []byte {
	p := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, path, p)
	if err != nil {
		logrus.Printf("DoPost: %s", path)
		logrus.Fatal(err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		logrus.Printf("DoPost: %s", path)
		logrus.Fatal(err)
	}
	defer resp.Body.Close()

	if (resp.Status != "200 OK") && (resp.Status != "201 Created") {
		logrus.Println("POST:", resp.Status, " path: ", path)
	} else {

		bodyResp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("DoPost:ReadAll: %s", path)
			log.Fatal(err)
		}
		return bodyResp
	}
	return nil

}

func (r *Rest_client) DoPatch(path string, body []byte, headers map[string]string) []byte {
	p := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPatch, path, p)

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := r.httpClient.Do(req)
	if err != nil {
		logrus.Printf("DoPatch: %s", path)
		logrus.Fatal(err)
	}
	defer resp.Body.Close()

	if (resp.Status != "200 OK") && (resp.Status != "201 Created") {
		logrus.Println("PATCH:", resp.Status, " path: ", path)
	} else {

		bodyResp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Printf("DoPatch:ReadAll: %s", path)
			logrus.Fatal(err)
		}
		return bodyResp
	}
	return nil

}

func (r *Rest_client) DoPut(path string, body io.Reader, headers map[string]string) []byte {

	req, err := http.NewRequest(http.MethodPut, path, body)

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := r.httpClient.Do(req)
	if err != nil {
		logrus.Printf("DoPut: %s", path)
		logrus.Fatal(err)
	}
	defer resp.Body.Close()

	if (resp.Status != "200 OK") && (resp.Status != "201 Created") {
		logrus.Println("PATCH:", resp.Status, " path: ", path)
	} else {

		bodyResp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Printf("DoPut:ReadAll: %s", path)
			logrus.Fatal(err)
		}
		return bodyResp
	}
	return nil

}
