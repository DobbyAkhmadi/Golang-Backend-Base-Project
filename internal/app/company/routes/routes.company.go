package routes

import (
	"backend/internal/app/company/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutesCompany Define your routes here
func SetupRoutesCompany(app *fiber.App) {
	// Group Company related APIs
	companyGroup := app.Group("/api/v1/company")
	companyGroup.Get("/", handlers.GetAllCompanies)
	companyGroup.Get("/:id", handlers.GetSingleCompany)
	companyGroup.Post("/", handlers.AddNewCompany)
	companyGroup.Put("/:id", handlers.UpdateCompany)
	companyGroup.Delete("/:id", handlers.DeleteCompany)

	// Group CompanyCategory related APIs
	companyCategoryGroup := app.Group("/api/v1/companyCategory")

	companyCategoryGroup.Get("/", handlers.GetAllCompanyCategories)
	companyCategoryGroup.Get("/:id", handlers.GetSingleCompanyCategory)
	companyCategoryGroup.Post("/", handlers.AddNewCompanyCategory)
	companyCategoryGroup.Put("/:id", handlers.UpdateCompanyCategory)
	companyCategoryGroup.Delete("/:id", handlers.DeleteCompanyCategory)
}
