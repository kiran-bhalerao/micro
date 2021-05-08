package publishers

import (
	"encoding/json"
	. "go-srv/pkg/redis"
)

type PostCreatedEventData struct {
	ID    float64 `json:"id"`
	Title string  `json:"title"`
}

type EventPublisher struct {
	subject string
	client  *PubSubClient
}

func (p EventPublisher) Publish(data PostCreatedEventData) {
	message, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	p.client.Publish(p.subject, string(message))
}

func NewPostCreatedPublisher(client *PubSubClient) *EventPublisher {
	return &EventPublisher{subject: "POST_CREATED", client: client}
}
