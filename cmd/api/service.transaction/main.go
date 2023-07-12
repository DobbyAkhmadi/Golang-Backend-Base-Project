package main

import (
	"backend/config"
	"backend/internal/app/transaction/routes"
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
	routes.SetupRoutesTransactionRoutes(app)

	// add more routes
	err = app.Listen(fmt.Sprintf(config.Config.GetString("TRANSACTION.SERVICE.HOST")+":%v",
		config.Config.GetString("TRANSACTION.SERVICE.PORT")))
	if err != nil {
		fmt.Println("Error starting Service transaction:", err)
		log.Fatal(err)
	}
}
