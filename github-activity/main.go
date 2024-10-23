package main

import (
	"encoding/json"
	"ncquang/github-activity/events"

	"fmt"
	"io"
	"log"
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

	page := 1

	for {
		data := makeRequest(client, userName, page)

		var listEvent []events.GitHubEventResponse

		err := json.Unmarshal(data, &listEvent)
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		for _, event := range listEvent {
			appEvent := events.CreateEventFactory(&event)
			if appEvent != nil {
				fmt.Printf("%v\n", appEvent)
			}
		}

		var userInput string
		for {
			fmt.Printf("Show more ? [y/n]\n")
			fmt.Scan(&userInput)

			if userInput == "y" {
				page = page + 1
				break
			} else if userInput == "n" {
				return
			} else {
				continue
			}
		}

	}
}

func makeRequest(client *http.Client, userName string, page int) []byte {
	response, err := client.Get(fmt.Sprintf("%s/users/%s/events?page=%d&per_page=10", gitHubEndPoint, userName, page))

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

	return body

}
