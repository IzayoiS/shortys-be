package main

import (
	"log"
	"os"

	"shortsy/config"
	"shortsy/database"
	"shortsy/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	app := fiber.New()
	database.Connect()
	database.Seed(database.DB)
	routes.RouterApp(app)

	port := os.Getenv("PORT")
	log.Println("app berjalan broo")
	
	app.Listen(":" + port)
}
