package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

type config struct {
	Url, Username, Password string
}

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getConfigPath() string {
	u, err := user.Current()
	checkError(err)

	return u.HomeDir + "/.suliechat.json"
}

func getConfig() config {
	configFile := getConfigPath()
	result := config{}

	if isFileExist(configFile) {
		data, err := ioutil.ReadFile(configFile)
		checkError(err)

		err = json.Unmarshal(data, &result)
		checkError(err)
	}

	return result
}

func saveConfig() {
	url, username, password := os.Args[2], os.Args[3], os.Args[4]
	cfg := &config{Url: url, Username: username, Password: password}
	configFile := getConfigPath()
	bytes, err := json.Marshal(cfg)
	checkError(err)

	err = ioutil.WriteFile(configFile, bytes, 0644)
	checkError(err)
}
