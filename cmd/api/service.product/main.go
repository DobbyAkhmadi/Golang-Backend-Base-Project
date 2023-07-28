package main

import (
	"backend/config"
	"backend/internal/app/product/routes"
	"backend/pkg/utils"
	"backend/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"strconv"
	"time"
)

const idleTimeout = 5 * time.Second

func main() {
	// Connect to database
	_, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	readTimeoutSecondsCount, _ := strconv.Atoi(config.Config.GetString("PRODUCT.SERVICE.SERVER_READ_TIMEOUT"))

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName:     "PRODUCT SERVICE Version : " + config.Config.GetString("SERVER.VERSION"),
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		IdleTimeout: idleTimeout,
	})

	// Define rate limiting configuration
	limiterConfig := limiter.Config{
		Max:        10,              // Maximum number of requests allowed per duration
		Expiration: 1 * time.Minute, // Duration for the rate limit window
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Use client IP as the key for rate limiting
		},
	}

	// Use the rate limiting middleware
	app.Use(limiter.New(limiterConfig))

	// Initialize routes
	routes.SetupRoutesProduct(app)

	// Use logger
	app.Use(logger.New())

	addr := config.Config.GetString("PRODUCT.SERVICE.HOST") + ":" + config.Config.GetString("PRODUCT.SERVICE.PORT")

	// start server with graceful shutdown
	utils.StartServerWithGracefulShutdown(app, addr)

}
