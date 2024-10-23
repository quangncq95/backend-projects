package events

import "fmt"

type pushEvent struct {
	Event   *GitHubEventResponse
	Payload *PushEventPayload
}

type PushEventPayload struct {
	Commits []interface{} `json:"commits"`
}

func (event *pushEvent) String() string {
	return fmt.Sprintf("Push %d commit to repo %s", len(event.Payload.Commits), event.Event.Repo.Name)
}
