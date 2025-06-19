package main

import (
	"log"
	"os"

	"shortsy/database"
	"shortsy/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// config.LoadEnv()
	database.Connect()
	database.Seed(database.DB)
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173,https://shortsy-web.up.railway.app,https://shortsy-web.netlify.app", 
        AllowHeaders: "Origin, Content-Type",
    }))
	routes.RouterApp(app)
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := os.Getenv("PORT")
	log.Println("app berjalan broo")
	
	app.Listen(":" + port)
}
