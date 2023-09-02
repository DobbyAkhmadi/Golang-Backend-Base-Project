package service

import "backend/internal/app/user/models"

// UserAuthService represents a products for user-related operations.
type UserAuthService interface {
	Login(dto *models.AuthLoginRequestDTO) (*models.AuthLoginResponseDTO, error)
}
