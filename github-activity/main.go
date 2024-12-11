package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Event struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Repo struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
}

var username string

func main() {

	username = os.Args[1]

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read request body: %v", err)
	}

	var data []Event

	if res.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Fatalf("failed to unmarshar request body: %v", err)
		}
	}else if res.StatusCode == http.StatusNotFound {
		fmt.Printf("username not found: %s",res.Status)
		return
	}else{
		log.Fatalf("error fetching user data: %s", res.Status )
	}

	for _, value := range data {
		fmt.Printf("Event: %s\nRepository: %s\nRepo URL: %s\nCreated At: %v\n\n", value.Type, value.Repo.Name, value.Repo.URL, value.CreatedAt)
	}

}
