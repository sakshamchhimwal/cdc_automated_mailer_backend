package controllers

import (
	"cdc_mailer/config"
	"cdc_mailer/models"
	"cdc_mailer/services"
	"cdc_mailer/utils"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func AddNewUser(c *fiber.Ctx) error {
	var adminDetails AdminAddUser

	if bodyError := c.BodyParser(&adminDetails); bodyError != nil {
		c.JSON(config.ErrorMessage(fiber.StatusBadRequest, "Invalid request body"))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	defaultPassword, err := utils.HashPassword(os.Getenv("INIT_PASS"))

	if err != nil {
		c.JSON(config.ErrorMessage(fiber.StatusInternalServerError, "Internal Server Error"))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	newUser := models.User{
		Name:         adminDetails.UserName,
		EmailAddress: adminDetails.UserEmail,
		Password:     defaultPassword,
		HasSignedUp:  false,
	}

	result := services.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(config.ErrorMessage(fiber.StatusInternalServerError, "Internal Server Error"))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "New user added successfully",
		"data": fiber.Map{
			"userId": newUser.ID,
		},
	})
	return c.SendStatus(fiber.StatusOK)
}

func AddCompanies(c *fiber.Ctx) error {
	companies := c.BodyRaw()
	fmt.Println(companies)

	return c.SendStatus(fiber.StatusOK)
}
