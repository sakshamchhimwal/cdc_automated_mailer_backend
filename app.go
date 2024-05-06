package main

import (
	"cdc_mailer/router"
	"cdc_mailer/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// Loading .env file into the GO-App
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())
	services.ConnectToDB()
	router.SetupRoutes(app)

	log.Fatal(app.Listen("localhost:3000"))
}
