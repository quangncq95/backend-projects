package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"ncquang/github-activity/events"
	"net/http"
	"os"
)

const gitHubEndPoint = "https://api.github.com"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing github user name")
	}
	client := &http.Client{}
	userName := os.Args[1]

	response, err := client.Get(gitHubEndPoint + "/users" + "/" + userName + "/events")
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Request failed : %s", http.StatusText(response.StatusCode))
		os.Exit(1)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	var listEvent []events.GitHubEventResponse

	err = json.Unmarshal(body, &listEvent)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	for _, event := range listEvent {
		appEvent := events.CreateEventFactory(&event)
		if appEvent != nil {
			fmt.Printf("%v\n", appEvent)
		}
	}
}
