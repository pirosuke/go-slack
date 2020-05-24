package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
Message describes the format for posting message.
*/
type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	AsUser  bool   `json:"as_user"`
}

/*
PostMessage posts message to slack.
*/
func PostMessage(token string, channel string, message string) ([]byte, error) {

	postMessage := Message{
		Channel: channel,
		Text:    message,
		AsUser:  true,
	}

	postMessageJSON, _ := json.Marshal(postMessage)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", strings.NewReader(string(postMessageJSON)))
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
