package routes

import (
	"backend/internal/app/transaction/handlers"
	"backend/internal/app/transaction/repository"
	"backend/internal/app/transaction/service"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutesTransactionRoutes Define your routes here
func SetupRoutesTransactionRoutes(app *fiber.App) {
	// Create the handlers instances with their respective dependencies
	// Create the repository and service dependencies
	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(&transactionRepository)

	// Create the handlers instances with their respective dependencies
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	transactionGroup := app.Group("/api/v1/transaction")
	transactionGroup.Get("/", transactionHandler.GetPaginationTransaction)
	transactionGroup.Post("/", transactionHandler.CreateTransaction)
	transactionGroup.Get("/:id", transactionHandler.GetTransactionById)
}
