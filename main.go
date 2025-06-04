package main

import (
	"os"

	"shortsy/config"
	"shortsy/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()
	app := fiber.New()
	database.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World GO!")
	})

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
