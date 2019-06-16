package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Date represents an imprecise date.
//type Date struct {
//	Year  int
//	Month time.Month
//	Day int
//
//}

type Message struct {
	ID        string
	Timestamp time.Time
}

func GetGuildMessages(guildId string) (result map[string]interface{}) {
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

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
