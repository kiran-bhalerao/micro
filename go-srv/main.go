package main

import (
	. "go-srv/events/listeners"
	. "go-srv/events/publishers"
	. "go-srv/pkg/redis"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	rdb := NewRedisClient()
	pubsub := PubSubClient{Client: rdb}

	go NewUserCreatedListener(&pubsub).Listen()

	app.Get("/", func(c *fiber.Ctx) error {
		NewPostCreatedPublisher(&pubsub).Publish(PostCreatedEventData{ID: 1, Title: "Post Title"})
		return c.JSON(fiber.Map{"msg": "Hello, World 👋!"})
	})

	app.Listen(":3000")
}
