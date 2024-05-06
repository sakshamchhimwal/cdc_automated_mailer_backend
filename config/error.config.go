package config

import "github.com/gofiber/fiber"

func ErrorMessage(status int, message string) fiber.Map {
	return fiber.Map{
		"status":  status,
		"message": message,
		"data":    nil,
	}
}
