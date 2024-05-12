package router

import (
	userControllers "cdc_mailer/controllers/user"
	fiber2 "github.com/gofiber/fiber/v2"
)

func SetupRoutesUser(user fiber2.Router) {
	user.Post("/updatePassword", userControllers.UpdatePassword)

	user.Get("/companies", userControllers.GetUserCompanies)
	user.Patch("/companies", userControllers.UpdateUserCompany)

	user.Post("/mail", userControllers.SendMailToCompany)

	user.Patch("/updateTemplate", userControllers.UpdateMailTemplate)
}
