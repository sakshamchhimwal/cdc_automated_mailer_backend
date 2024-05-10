package controllers

import (
	"cdc_mailer/models"
	"cdc_mailer/services"
	"cdc_mailer/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var userDetails LoginInput
	if bodyError := c.BodyParser(&userDetails); bodyError != nil {
		return utils.SetError(c, fiber.StatusBadRequest, "Invalid Request body")
	}

	emailAddress := userDetails.EmailAddress
	password := userDetails.Password

	var user models.User
	results := services.DB.Where("email_address = ?", emailAddress).First(&user)

	if results.RowsAffected == 0 {
		return utils.SetError(c, fiber.StatusUnauthorized, "User not found")
	}

	if !utils.ComparePasswords(password, user.Password) {
		return utils.SetError(c, fiber.StatusUnauthorized, "Passwords do not match")
	}

	payload := utils.Payload{Id: uint64(user.ID)}

	token, jwtError := utils.ProvideJWT(&payload)

	if jwtError != nil {
		return utils.SetError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	_ = c.JSON(fiber.Map{
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
