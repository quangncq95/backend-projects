package events

type GitHubRepo struct {
	ID   int
	Name string
	Url  string
}

type GitHubEventType string
type GitHubEventResponse struct {
	ID      string
	Type    GitHubEventType
	Repo    GitHubRepo
	Payload any
}

type Event interface {
	String() string
}

const (
	PushEvent        GitHubEventType = "PushEvent"
	CreateEvent      GitHubEventType = "CreateEvent"
	PullRequestEvent GitHubEventType = "PullRequestEvent"
)

func CreateEventFactory(res *GitHubEventResponse) Event {
	switch res.Type {
	case PushEvent:
		m := res.Payload.(map[string]interface{})
		if payload, ok := m["commits"].([]interface{}); ok {
			return &pushEvent{
				Event:   res,
				Payload: &PushEventPayload{Commits: payload},
			}
		}

		return nil
	case CreateEvent:
		return &createEvent{
			Event: res,
		}
	case PullRequestEvent:
		m := res.Payload.(map[string]interface{})
		return &pullRequestEvent{
			Event: res,
			Payload: &pullRequestEventPayload{
				action: m["action"].(string),
				number: m["number"].(float64),
			},
		}
	}

	return nil
}
