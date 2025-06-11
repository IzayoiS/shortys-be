package controllers

import (
	"os"
	service "shortsy/service"
	"github.com/gofiber/fiber/v2"
)

func PostLink(c *fiber.Ctx) error {
	var body struct {
		OriginalURL string `json:"original_url"`
	}

	port := os.Getenv("PORT")

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	url, err := service.PostLink(body.OriginalURL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	port := os.Getenv("PORT")
	return c.JSON(fiber.Map{
		"message": "Successfully shortener link",
		"short": "http://127.0.0.1:"+ port + "/" + url.ShortCode,
	})
}

func GetLink(c *fiber.Ctx) error {
	shortCode := c.Params("shortCode")
	shortlink, err := service.GetLink(shortCode)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "ShortLink not found",
		})
	}

	return c.Redirect(shortlink.OriginalURL,fiber.StatusTemporaryRedirect)
}
