package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
	User      string `json:"user,omitempty"`
	UserName  string `json:"username,omitempty"`
}

type MessageResponse struct {
	Version       string    `json:"appVersion"`
	Data          []Message `json:"data"`
	Notifications []byte    `json:"-"`
	Success       bool      `json:"success"`
	UserV         int       `json:"-"`
}

func GetGuildMessages(guildId string) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := fmt.Sprintf("https://habitica.com/api/v3/groups/%s/chat", guildId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("x-api-user", "INSERT_API_ID_HERE")
	req.Header.Add("x-api-key", "INSERT_API_KEY_HERE")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var result MessageResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result.Data[0])
	return
}
