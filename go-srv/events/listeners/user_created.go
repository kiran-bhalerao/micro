package listeners

import (
	"fmt"
	. "go-srv/pkg/redis"
)

type EventListener struct {
	subject string
	client  *PubSubClient
}

type UserCreatedEventData struct {
	id   float64
	name string
}

// Cast method basically map event pyaload to a UserCreatedEvent struct
func (u EventListener) Cast(payload map[string]interface{}) UserCreatedEventData {
	return UserCreatedEventData{
		id:   payload["id"].(float64),
		name: payload["name"].(string),
	}
}

func (u EventListener) Listen() {
	_, listener := u.client.Subscribe(u.subject)

	listener.Onmessage(func(payload map[string]interface{}) {
		event := u.Cast(payload)

		fmt.Println(event.id, event.name)
	})
}

func NewUserCreatedListener(client *PubSubClient) *EventListener {
	return &EventListener{subject: "USER_CREATED", client: client}
}
