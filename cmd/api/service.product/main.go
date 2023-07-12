package main

import (
	"backend/config"
	"backend/internal/app/product/routes"
	"backend/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const idleTimeout = 5 * time.Second

func main() {
	// Connect to database
	_, err := utils.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	// Use logger
	app.Use(logger.New())

	// Initialize routes
	routes.SetupRoutesProduct(app)

	err = app.Listen(fmt.Sprintf(config.Config.GetString("PRODUCT.SERVICE.HOST")+":%v",
		config.Config.GetString("PRODUCT.SERVICE.PORT")))
	if err != nil {
		fmt.Println("Error starting Service product:", err)
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt	 or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	//connection.Close()
	fmt.Println("Fiber was successful shutdown.")
}
