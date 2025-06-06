package controllers

import (
	"github.com/gofiber/fiber/v2"
	service "shortsy/service"
)

func PostLink(c *fiber.Ctx) error {
	var body struct {
		OriginalURL string `json:"original_url"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	url,err := service.PostLink(body.OriginalURL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"url": fiber.Map{
			"id":    url.Id,
			"original_url":  url.OriginalURL,
			"short": url.ShortCode,
		},
	})
}
