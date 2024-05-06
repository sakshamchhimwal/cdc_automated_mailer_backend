package controllers

import (
	"cdc_mailer/config"
	"cdc_mailer/models"
	"cdc_mailer/services"
	"cdc_mailer/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UpdatePassword(c *fiber.Ctx) error {

	var passowrdDetails PasswordInput
	if bodyError := c.BodyParser(&passowrdDetails); bodyError != nil {
		c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid Request Format",
			"data":    nil,
		})
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newPassword := passowrdDetails.NewPassword
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(string)

	var findUser models.User
	result := services.DB.First(&findUser, userId)

	if result.RowsAffected == 0 {
		c.JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "User not found",
			"data":    nil,
		})
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	newHashedPassword, err := utils.HashPassword(newPassword)

	if err != nil {
		c.JSON(config.ErrorMessage(fiber.StatusInternalServerError, "Internal Server Error"))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	findUser.Password = newHashedPassword

	services.DB.Save(&findUser)

	c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Password Update Successful",
	})
	return c.SendStatus(fiber.StatusOK)
}
