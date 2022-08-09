package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"data": "get de usuarios",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server on port 3000")
}
