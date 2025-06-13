package main

import (
	"log"
	"os"

	"shortsy/database"
	"shortsy/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.Connect()
	database.Seed(database.DB)
	routes.RouterApp(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := os.Getenv("PORT")
	log.Println("app berjalan broo")
	
	app.Listen(":" + port)
}
