package events

import "fmt"

type watchEvent struct {
	Event   *GitHubEventResponse
	Payload *watchEventPayload
}

type watchEventPayload struct {
	action string
}

func (event *watchEvent) String() string {
	return fmt.Sprintf("%s repo %s", event.Payload.action, event.Event.Repo.Name)
}
