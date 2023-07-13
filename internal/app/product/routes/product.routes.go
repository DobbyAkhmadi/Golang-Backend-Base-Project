package routes

import (
	"backend/internal/app/product/handlers"
	"backend/internal/app/product/repository"
	"backend/internal/app/product/service"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutesProduct Define your routes here
func SetupRoutesProduct(app *fiber.App) {
	// Create the handlers instances with their respective dependencies
	// Create the repository and service dependencies
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(&productRepository)

	// Create the handlers instances with their respective dependencies
	productHandler := handlers.NewProductHandler(productService)

	productGroup := app.Group("/api/v1/product")
	productGroup.Get("/", productHandler.GetPaginationProduct)
	productGroup.Get("/:id", productHandler.GetProductByID)
	productGroup.Post("/", productHandler.CreateNewProduct)
	productGroup.Put("/:id", productHandler.UpdateExistingProduct)
	productGroup.Delete("/:id", productHandler.DeleteProductByID)

}
