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
AppConfigDirName は設定が保存されるディレクトリ名です。ホームディレクトリ下に指定された名前のディレクトリが作成されます。
*/
const AppConfigDirName = "go-slack"

/*
AppConfig はアプリケーション設定情報定義です。
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
