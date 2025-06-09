package routes

import (
	controller "shortsy/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouterApp(c*fiber.App){
	c.Post("/",controller.PostLink)
	c.Get("/:shortCode",controller.GetLink)
}
