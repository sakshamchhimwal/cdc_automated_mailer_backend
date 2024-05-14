package router

import (
	apiControllers "cdc_mailer/controllers/api"
	authControllers "cdc_mailer/controllers/auth"
	"cdc_mailer/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Test Route
	api := app.Group("/api", logger.New())
	api.Get("/", apiControllers.Hello)

	// Auth Router
	auth := api.Group("/auth", logger.New())
	auth.Post("/login", authControllers.Login)

	// Admin Routes
	admin := api.Group("/admin", logger.New())
	SetupRoutesAdmin(admin)

	// User Routes
	user := api.Group("/user", middleware.Protected(), logger.New())
	SetupRoutesUser(user)
}
