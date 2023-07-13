package handlers

import (
	"backend/internal/app/user/service"
	"github.com/gofiber/fiber/v2"
)

// UserHandler represents a handler for user-related operations.
type UserHandler struct {
	userService     service.UserService
	userAuthService service.UserAuthService
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userService service.UserService, userAuthService service.UserAuthService) *UserHandler {
	return &UserHandler{
		userService:     userService,
		userAuthService: userAuthService,
	}
}

func (h *UserHandler) LoginUser(ctx *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) RefreshToken(ctx *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) CreateNewUser(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) UpdateExistingUser(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) DeleteUserByID(ctx *fiber.Ctx) error {
	return nil
}
