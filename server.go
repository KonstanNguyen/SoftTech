package main

import (
	"SoftTech/blogs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	blogs.Routes(app)
	log.Fatal(app.Listen(":4000"))
}
