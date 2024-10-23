package events

import "fmt"

type createEvent struct {
	Event *GitHubEventResponse
}

func (event *createEvent) String() string {
	return fmt.Sprintf("Create new repo : %v", event.Event.Repo.Name)
}
