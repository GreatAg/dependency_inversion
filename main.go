package main

import "github.com/gofiber/fiber/v2"

var (
	app = fiber.New()
)

func main() {
	init_routes()
	app.Listen(":8080")
}
