package main

import (
	"github.com/adiga27/web-hosting-go/internals/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	routes.ManualRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Page not found!",
		})
	})

	app.Listen(":8080")
}
