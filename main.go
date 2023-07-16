package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		time.Sleep(time.Microsecond * 10)
		return c.SendString("Hello World")
	})

	app.Listen("localhost:8000")
}
