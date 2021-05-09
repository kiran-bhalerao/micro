package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	REDIS_HOST := os.Getenv("REDIS_HOST")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", REDIS_HOST),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

/************ */
type ChannelMessage struct {
	Channel <-chan *redis.Message
}

func (l *ChannelMessage) Onmessage(fn func(payload map[string]interface{})) {
	for msg := range l.Channel {
		var data map[string]interface{}
		json.Unmarshal([]byte(msg.Payload), &data)

		fn(data)
	}
}

type PubSubClient struct {
	Client *redis.Client
}

/************ SUB */
func (ps *PubSubClient) Subscribe(channelId string) (<-chan *redis.Message, *ChannelMessage) {
	pubsub := ps.Client.Subscribe(ctx, channelId)

	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	ch := pubsub.Channel()

	return ch, &ChannelMessage{Channel: ch}
}

/************ PUB */
func (ps *PubSubClient) Publish(channelId string, message string) {
	err := ps.Client.Publish(ctx, channelId, message).Err()
	if err != nil {
		panic(err)
	}
}
