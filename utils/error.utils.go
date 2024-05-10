package utils

import "github.com/gofiber/fiber/v2"

func errorMessage(status int, message string) fiber.Map {
	return fiber.Map{
		"status":  status,
		"message": message,
		"data":    nil,
	}
}

func SetError(ctx *fiber.Ctx, status int, message string) error {
	_ = ctx.JSON(errorMessage(status, message))
	return ctx.SendStatus(status)
}
