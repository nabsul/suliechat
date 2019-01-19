package main

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
)

func request(method, path string, data io.Reader) string {
	cfg := getConfig()
	root := "http://localhost:7071/api/"
	if cfg.Url != "localhost" {
		root = "https://" + cfg.Url + ".azurewebsites.net/api/"
	}
	url := root + path

	client := &http.Client{}
	req, err := http.NewRequest(method, url, data)
	checkError(err)

	auth := base64.StdEncoding.EncodeToString([]byte(cfg.Username + ":" + cfg.Password))
	req.Header.Add("Authorization", "Basic " + auth)
	req.Header.Add("Content-Type", "application/json")

	result, err := client.Do(req)
	checkError(err)

	body, err := ioutil.ReadAll(result.Body)
	checkError(err)

	return string(body)
}

