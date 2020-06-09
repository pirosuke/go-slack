package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pirosuke/go-slack"
)

/*
AppConfigDirName is the directory name to save settings.
*/
const AppConfigDirName = "go-slack"

/*
AppConfig describes setting format.
*/
type AppConfig struct {
	Token string `json:"token"`
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func main() {
	postCmd := flag.NewFlagSet("post", flag.ExitOnError)
	workspace := postCmd.String("w", "", "Workspace Name")
	channel := postCmd.String("c", "#general", "Channel")
	message := postCmd.String("m", "", "Message To Post")
	postCmd.Parse(os.Args[2:])

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	appConfigDirPath := filepath.Join(userConfigDir, AppConfigDirName)
	os.MkdirAll(appConfigDirPath, os.ModePerm)

	appConfigFilePath := filepath.Join(appConfigDirPath, *workspace+".json")
	if !fileExists(appConfigFilePath) {
		fmt.Println("config file does not exist: " + appConfigFilePath)
		return
	}

	jsonContent, err := ioutil.ReadFile(appConfigFilePath)
	if err != nil {
		fmt.Println("failed to read config file: " + appConfigFilePath)
		return
	}

	appConfig := new(AppConfig)
	if err := json.Unmarshal(jsonContent, appConfig); err != nil {
		fmt.Println("failed to read config file: " + appConfigFilePath)
		return
	}

	slack.PostMessage(appConfig.Token, *channel, *message)
}
