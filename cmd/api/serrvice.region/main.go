package main

import (
	"backend/config"
	"backend/internal/app/REGION/routes"
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

	readTimeoutSecondsCount, _ := strconv.Atoi(config.Config.GetString("REGION.SERVICE.SERVER_READ_TIMEOUT"))

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName:     "REGION SERVICE Version : " + config.Config.GetString("SERVER.VERSION"),
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		IdleTimeout: idleTimeout,
	})

	// Initialize routes
	routes.SetupRoutesRegion(app)

	// Use logger
	app.Use(logger.New())

	addr := config.Config.GetString("REGION.SERVICE.HOST") + ":" + config.Config.GetString("REGION.SERVICE.PORT")

	// start server with graceful shutdown
	utils.StartServerWithGracefulShutdown(app, addr)

}
