package events

import "fmt"

type pullRequestEvent struct {
	Event   *GitHubEventResponse
	Payload *pullRequestEventPayload
}

type pullRequestEventPayload struct {
	action string
	number float64
}

func (event *pullRequestEvent) String() string {
	return fmt.Sprintf("%s pull request number %d repo:%s", event.Payload.action, int(event.Payload.number), event.Event.Repo.Name)
}
