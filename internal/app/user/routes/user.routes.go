package routes

import (
	"backend/internal/app/user/handlers"
	"backend/internal/app/user/repository"
	"backend/internal/app/user/service"
	"backend/pkg/jwt"
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
	userGroup.Post("/auth/login", userHandler.LoginUser)
	//userGroup.Put("/auth/refresh", userHandler.RefreshToken)
	//userGroup.Put("/auth/logout", userHandler.RefreshToken)
	//userGroup.Put("/auth/verify", userHandler.RefreshToken)

	userGroup.Get("/", userHandler.GetPaginationUser)
	userGroup.Get("/:id", userHandler.GetUserByID)
	userGroup.Post("/", jwt.Auth(), userHandler.CreateNewUser)
	userGroup.Put("/:id", userHandler.UpdateExistingUser)
	userGroup.Delete("/:id", userHandler.DeleteUserByID)

}
