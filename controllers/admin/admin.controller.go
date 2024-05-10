package controllers

import (
	"cdc_mailer/models"
	"cdc_mailer/services"
	"cdc_mailer/utils"
	"github.com/gofiber/fiber/v2"
	"os"
)

func AddNewUser(c *fiber.Ctx) error {
	var adminDetails AdminAddUser

	if bodyError := c.BodyParser(&adminDetails); bodyError != nil {
		return utils.SetError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	defaultPassword, err := utils.HashPassword(os.Getenv("INIT_PASS"))

	if err != nil {
		return utils.SetError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	newUser := models.User{
		Name:         adminDetails.UserName,
		EmailAddress: adminDetails.UserEmail,
		Password:     defaultPassword,
		HasSignedUp:  false,
	}

	result := services.DB.Create(&newUser)

	if result.Error != nil {
		return utils.SetError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	_ = c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "New user added successfully",
		"data": fiber.Map{
			"userId": newUser.ID,
		},
	})
	return c.SendStatus(fiber.StatusOK)
}

func AddCompanies(c *fiber.Ctx) error {

	var companyData []models.Company

	if bodyError := c.BodyParser(&companyData); bodyError != nil {
		return utils.SetError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result := services.DB.Create(&companyData)

	if result.Error != nil {
		return utils.SetError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	var companiesID []uint
	for _, companies := range companyData {
		companiesID = append(companiesID, companies.ID)
	}

	_ = c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "New companies added successfully",
		"data": fiber.Map{
			"companiesId": companiesID,
		},
	})
	return c.SendStatus(fiber.StatusOK)
}
