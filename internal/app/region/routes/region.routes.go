package routes

import (
	"backend/internal/app/region/handlers"
	"backend/internal/app/region/repository"
	"backend/internal/app/region/service"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutesRegion Define your routes here
func SetupRoutesRegion(app *fiber.App) {
	// Create the handlers instances with their respective dependencies
	// Create the repository and service dependencies
	regionRepository := repository.NewRegionRepository()
	regionService := service.NewRegionService(&regionRepository)

	// Create the handlers instances with their respective dependencies
	regionHandler := handlers.NewRegionHandler(regionService)

	regionGroup := app.Group("/api/v1/region")
	regionGroup.Get("/village", regionHandler.GetPaginationVillage)
	regionGroup.Get("/province", regionHandler.GetPaginationVillage)
	regionGroup.Get("/regency", regionHandler.GetPaginationVillage)
	regionGroup.Get("/district", regionHandler.GetPaginationVillage)
}
