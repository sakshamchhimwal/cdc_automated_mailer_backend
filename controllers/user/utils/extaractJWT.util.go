package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

func GetUserId(c *fiber.Ctx) uint {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId, _ := strconv.ParseUint(claims["userId"].(string), 10, 32)
	return uint(userId)
}
