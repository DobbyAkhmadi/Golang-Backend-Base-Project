package routes

import (
	"backend/internal/app/user/handlers"
	"backend/internal/app/user/repository"
	"backend/internal/app/user/service"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutesUser Define your routes here
func SetupRoutesUser(app *fiber.App) {
	// Create the handlers instances with their respective dependencies
	// Create the repository and service dependencies
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(&userRepository)
	authService := service.NewUserAuthService(&userRepository)

	// Create the handlers instances with their respective dependencies
	userHandler := handlers.NewUserHandler(userService, authService)

	userGroup := app.Group("/api/v1/user")
	userGroup.Get("/", userHandler.GetPaginationUser)
	userGroup.Get("/:id", userHandler.GetUserByID)
	userGroup.Post("/", userHandler.CreateNewUser)
	userGroup.Put("/:id", userHandler.UpdateExistingUser)
	userGroup.Delete("/:id", userHandler.DeleteUserByID)
}
