package router

import (
	adminControllers "cdc_mailer/controllers/admin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"os"
)

func SetupRoutesAdmin(admin fiber.Router) {
	adminName := os.Getenv("ADMIN_NAME")
	adminKey := os.Getenv("ADMIN_KEY")
	admin.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			adminName: adminKey,
		},
	}))
	admin.Post("/newUser", adminControllers.AddNewUser)
	admin.Post("/newCompanies", adminControllers.AddCompanies)
}
