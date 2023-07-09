package main

import (
	"backend/internal/app/company/handlers"
	"backend/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {

	// Connect to database
	_, err := entity.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Response with a hello message for calling root path
	app.Get("/", hello)

	// Use logger
	app.Use(logger.New())

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

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
