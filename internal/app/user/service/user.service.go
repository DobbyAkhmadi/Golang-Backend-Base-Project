package service

import (
	models2 "backend/internal/app/user/models"
	"backend/pkg/utils"
)

// UserService represents a service for user-related operations.
type UserService interface {
	Create(request *models2.CreateUserRequestDTO) (models2.GetUserResponseDTO, error)
	Update(id string, request *models2.UpdateUserRequestDTO) (models2.GetUserResponseDTO, error)
	Delete(id string) (models2.GetUserResponseDTO, error)
	GetPagination(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
	GetUserByID(id string) (models2.GetUserResponseDTO, error)
}
