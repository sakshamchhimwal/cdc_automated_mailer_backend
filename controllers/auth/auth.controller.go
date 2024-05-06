package controllers

import (
	"cdc_mailer/config"
	"cdc_mailer/models"
	"cdc_mailer/services"
	"cdc_mailer/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var userDetails LoginInput
	if bodyError := c.BodyParser(&userDetails); bodyError != nil {
		return c.JSON(config.ErrorMessage(fiber.StatusBadRequest, "Invalid request body"))
	}

	emailAddress := userDetails.EmailAddress
	password := userDetails.Password

	var user models.User
	results := services.DB.Where("email_address = ?", emailAddress).First(&user)

	if results.RowsAffected == 0 {
		c.JSON(config.ErrorMessage(fiber.StatusUnauthorized, "User not found"))
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !utils.ComparePasswords(password, user.Password) {
		c.JSON(config.ErrorMessage(fiber.StatusUnauthorized, "Passwords do not match"))
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	payload := utils.Payload{Id: uint64(user.ID)}

	token, jwtError := utils.ProvideJWT(&payload)

	if jwtError != nil {
		fmt.Println(jwtError)
		c.JSON(config.ErrorMessage(fiber.StatusInternalServerError, "Internal Server Error"))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "User login successful",
		"data": fiber.Map{
			"emailAddress": user.EmailAddress,
			"name":         user.Name,
			"userSignUp":   user.HasSignedUp,
		},
		"token": token,
	})
	return c.SendStatus(fiber.StatusOK)
}
