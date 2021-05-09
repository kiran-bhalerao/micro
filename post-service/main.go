package main

import (
	. "post-service/events/listeners"
	. "post-service/events/publishers"
	. "post-service/pkg/redis"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	rdb := NewRedisClient()
	pubsub := PubSubClient{Client: rdb}

	go NewUserCreatedListener(&pubsub).Listen()

	app.Get("/api/post", func(c *fiber.Ctx) error {
		NewPostCreatedPublisher(&pubsub).Publish(PostCreatedEventData{ID: 1, Title: "Post Title"})
		return c.JSON(fiber.Map{"msg": "Hello, Post Service ❤️!"})
	})

	app.Listen(":3000")
}
