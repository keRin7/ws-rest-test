package rest_client

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type Rest_client struct {
	config     *Config
	httpClient *http.Client
}

func NewRestClient(config *Config) *Rest_client {
	return &Rest_client{
		config:     config,
		httpClient: &http.Client{},
	}
}

func (r *Rest_client) DoGet(path string, headers map[string]string) []byte {

	//client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, r.config.url+path, nil)

	for k, v := range headers {
		//fmt.Printf("%s %s", k, v)
		req.Header.Add(k, v)
	}
	resp, err := r.httpClient.Do(req)
	if err != nil {
		log.Printf("DoGet: %s", path)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("DoGet:ReadAll: %s", path)
		log.Fatal(err)
	}

	return body
	/* for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	} */
}

func (r *Rest_client) DoPost(path string, b []byte) []byte {
	p := bytes.NewReader(b)
	resp, err := http.Post(r.config.url+path, "application/json", p)
	if err != nil {
		log.Printf("DoPost: %s", path)
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("DoPost:ReadAll: %s", path)
		log.Fatal(err)
	}

	//fmt.Println(string(body))

	return body
}

func (r *Rest_client) DoPatch(path string, b []byte, headers map[string]string) []byte {
	p := bytes.NewReader(b)
	//resp, err := http.

	req, err := http.NewRequest(http.MethodPatch, r.config.url+path, p)

	for k, v := range headers {
		//fmt.Printf("%s %s", k, v)
		req.Header.Add(k, v)
	}
	resp, err := r.httpClient.Do(req)
	if err != nil {
		log.Printf("DoPatch: %s", path)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("DoPatch:ReadAll: %s", path)
		log.Fatal(err)
	}

	//fmt.Println(string(body))

	return body
}

/* data := []byte(`{"foo":"bar"}`)
r := bytes.NewReader(data)
resp, err := http.Post("http://example.com/upload", "application/json", r)
if err != nil {
    return err
} */
