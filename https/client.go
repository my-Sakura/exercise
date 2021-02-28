package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	url = "https://127.0.0.1:1234"
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func BcjClient(request []string) ([]bool, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	var req struct {
		Request []string `json: "request"`
	}
	req.Request = request
	byteData, err := json.Marshal(req)

	reader := bytes.NewReader(byteData)
	r, err := http.NewRequest("GET", url, reader)
	r.Header.Set("Content-Type", "application/json")
	checkError(err)

	resp, err := client.Do(r)
	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	var Resp struct {
		Response []bool `json: "response"`
	}
	json.Unmarshal(body, &Resp)

	return Resp.Response, nil
}
