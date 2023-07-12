package main

import (
	"backend/config"
	"backend/internal/app/transaction/routes"
	"backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"strconv"
	"time"
)

const idleTimeout = 5 * time.Second

func main() {
	// Connect to database
	_, err := utils.Connect()
	if err != nil {
		log.Fatal(err)
	}

	readTimeoutSecondsCount, _ := strconv.Atoi(config.Config.GetString("TRANSACTION.SERVICE.SERVER_READ_TIMEOUT"))

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName:     "TRANSACTION SERVICE Version : " + config.Config.GetString("SERVER.VERSION"),
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		IdleTimeout: idleTimeout,
	})

	// Initialize routes
	routes.SetupRoutesTransactionRoutes(app)

	// Use logger
	app.Use(logger.New())

	addr := config.Config.GetString("TRANSACTION.SERVICE.HOST") + ":" + config.Config.GetString("TRANSACTION.SERVICE.PORT")

	// start server with graceful shutdown
	utils.StartServerWithGracefulShutdown(app, addr)
}