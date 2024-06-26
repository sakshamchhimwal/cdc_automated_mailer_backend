package middleware

import (
	"cdc_mailer/utils"
	"fmt"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() func(*fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")
	return jwtware.New(
		jwtware.Config{
			SigningKey:   jwtware.SigningKey{Key: []byte(secret)},
			ErrorHandler: jwtError,
		},
	)
}

func jwtError(c *fiber.Ctx, err error) error {
	fmt.Println(err.Error())
	if err.Error() == "Missing or malformed JWT" {
		return utils.SetError(c, fiber.StatusUnauthorized, "Missing or malformed JWT")

	} else {
		return utils.SetError(c, fiber.StatusUnauthorized, "Invalid or expired JWT")
	}
}
