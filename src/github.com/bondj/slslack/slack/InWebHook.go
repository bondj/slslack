package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Text       string `json:"text"`
	Username   string `json:"username"`
	Icon_url   string `json:"icon_url"`
	Icon_emoji string `json:"icon_emoji"`
	Channel    string `json:"channel"`
}

func SendMessage(client *http.Client, m Message, target string) {

	b, err := json.Marshal(m)

	fmt.Println(b)

	req, err := http.NewRequest("POST", target, bytes.NewBuffer(b))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
