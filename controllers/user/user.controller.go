package controllers

import (
	userUtils "cdc_mailer/controllers/user/utils"
	"cdc_mailer/models"
	"cdc_mailer/services"
	"cdc_mailer/utils"
	"github.com/gofiber/fiber/v2"
)

func UpdatePassword(c *fiber.Ctx) error {

	var passwordDetails PasswordInput
	if bodyError := c.BodyParser(&passwordDetails); bodyError != nil {
		return utils.SetError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	newPassword := passwordDetails.NewPassword
	userId := userUtils.GetUserId(c)

	var findUser models.User
	result := services.DB.First(&findUser, uint(userId))

	if result.RowsAffected == 0 {
		return utils.SetError(c, fiber.StatusUnauthorized, "No User Found")
	}

	newHashedPassword, err := utils.HashPassword(newPassword)

	if err != nil {
		return utils.SetError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	findUser.Password = newHashedPassword

	services.DB.Save(&findUser)

	_ = c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Password Update Successful",
	})
	return c.SendStatus(fiber.StatusOK)
}

func GetUserCompanies(c *fiber.Ctx) error {
	userId := userUtils.GetUserId(c)

	if userUtils.VerifyUser(userId) != nil {
		return utils.SetError(c, fiber.StatusUnauthorized, "No User Found")
	}

	result := services.DB.Model(&models.Company{}).Where("HandlerID = ?", userId)

	if result.Error != nil {
		return utils.SetError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	_ = c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"data":    result,
		"message": "User Companies Successful",
	})
	return c.SendStatus(fiber.StatusOK)
}

func UpdateUserCompany(c *fiber.Ctx) error {
	userId := userUtils.GetUserId(c)
	var companyUpdateDetails UpdateCompanyInput

	bodyError := c.BodyParser(&companyUpdateDetails)
	if bodyError != nil {
		return utils.SetError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if userUtils.VerifyUser(userId) != nil {
		return utils.SetError(c, fiber.StatusUnauthorized, "No User Found")
	}

	findCompany, findCompanyErr := userUtils.VerifyUserCompany(userId, companyUpdateDetails.CompanyId)

	if findCompanyErr != nil {
		return utils.SetError(c, fiber.StatusUnauthorized, "No Company Found")
	}

	if findCompany.MailVerified == false && companyUpdateDetails.MailVerified == true {
		findCompany.MailVerified = companyUpdateDetails.MailVerified
	}

	services.DB.Save(&findCompany)

	_ = c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Password Update Successful",
		"data":    findCompany,
	})

	return c.SendStatus(fiber.StatusOK)
}

func UpdateMailTemplate(c *fiber.Ctx) error {
	userId := userUtils.GetUserId(c)
	var companyTemplateUpdateDetails CompanyTemplateUpdateDetails

	bodyError := c.BodyParser(&companyTemplateUpdateDetails)
	if bodyError != nil {
		return utils.SetError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if userUtils.VerifyUser(userId) != nil {
		return utils.SetError(c, fiber.StatusUnauthorized, "No User Found")
	}

	findCompany, findCompanyErr := userUtils.VerifyUserCompany(userId, companyTemplateUpdateDetails.CompanyId)

	if findCompanyErr != nil {
		return utils.SetError(c, fiber.StatusUnauthorized, "No Company Found")
	}

	if findCompany.MailSent == true {
		return utils.SetError(c, fiber.StatusForbidden, "Mail already sent")
	}

	if companyTemplateUpdateDetails.TemplateNumber == 1 {
		findCompany.Template1 = companyTemplateUpdateDetails.TemplateContent
	} else if companyTemplateUpdateDetails.TemplateNumber == 2 {
		findCompany.Template2 = companyTemplateUpdateDetails.TemplateContent
	} else if companyTemplateUpdateDetails.TemplateNumber == 3 {
		findCompany.Template3 = companyTemplateUpdateDetails.TemplateContent
	} else {
		return utils.SetError(c, fiber.StatusNotFound, "Template does not exist")
	}

	_ = c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Template Update Successful",
	})
	return c.SendStatus(fiber.StatusOK)

}

func SendMailToCompany(c *fiber.Ctx) error {
	userId := userUtils.GetUserId(c)
	var companyMailingDetails CompanyMailingDetails

	bodyError := c.BodyParser(&companyMailingDetails)
	if bodyError != nil {
		return utils.SetError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if userUtils.VerifyUser(userId) != nil {
		return utils.SetError(c, fiber.StatusUnauthorized, "No User Found")
	}

	findCompany, findCompanyErr := userUtils.VerifyUserCompany(userId, companyMailingDetails.CompanyId)

	if findCompanyErr != nil {
		return utils.SetError(c, fiber.StatusUnauthorized, "No Company Found")
	}

	if findCompany.MailSent == true {
		return utils.SetError(c, fiber.StatusForbidden, "Mail already sent")
	}

	var mailTemplate string

	if companyMailingDetails.TemplateNumber == 1 {
		mailTemplate = findCompany.Template1
	} else if companyMailingDetails.TemplateNumber == 2 {
		mailTemplate = findCompany.Template2
	} else if companyMailingDetails.TemplateNumber == 3 {
		mailTemplate = findCompany.Template3
	} else {
		return utils.SetError(c, fiber.StatusNotFound, "No Template Exists")
	}

	if len(mailTemplate) == 0 {
		return utils.SetError(c, fiber.StatusForbidden, "Cannot send empty mail")
	}

	mailError := services.SendMail(companyMailingDetails.CompanyId, mailTemplate, findCompany.HrEmail)
	if mailError != nil {
		return utils.SetError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	_ = c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Mail Send Successful",
	})
	return c.SendStatus(fiber.StatusOK)

}
