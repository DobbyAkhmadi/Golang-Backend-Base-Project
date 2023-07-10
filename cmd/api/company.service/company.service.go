package main

import (
	"backend/config"
	"backend/internal/app/company/routes"
	"backend/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	// Connect to database
	_, err := utils.Connect()
	if err != nil {
		log.Fatal(err)
	}
	// Create a new Fiber instance
	app := fiber.New()

	// Use logger
	app.Use(logger.New())

	// Initialize routes
	routes.SetupRoutesCompany(app)

	// add more routes

	err = app.Listen(fmt.Sprintf(config.Config.GetString("COMPANY.SERVICE.HOST")+":%v",
		config.Config.GetString("COMPANY.SERVICE.PORT")))
	if err != nil {
		log.Fatal(err)
	}
}
