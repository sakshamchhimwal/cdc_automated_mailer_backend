package router

import (
	adminControllers "cdc_mailer/controllers/admin"
	apiControllers "cdc_mailer/controllers/api"
	authControllers "cdc_mailer/controllers/auth"
	userControllers "cdc_mailer/controllers/user"
	"cdc_mailer/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	// Admin Route
	adminName := os.Getenv("ADMIN_NAME")
	adminKey := os.Getenv("ADMIN_KEY")

	admin := app.Group("/admin", logger.New())
	admin.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			adminName: adminKey,
		},
	}))
	admin.Post("/newUser", adminControllers.AddNewUser)
	admin.Post("/newCompanies", adminControllers.AddCompanies)

	// Test Route
	api := app.Group("/api", logger.New())
	api.Get("/", apiControllers.Hello)

	// Auth Router
	auth := app.Group("/auth", logger.New())
	auth.Post("/login", authControllers.Login)

	// User Routes
	user := api.Group("/user", middleware.Protected(), logger.New())
	user.Post("/updatePassword", userControllers.UpdatePassword)
}
