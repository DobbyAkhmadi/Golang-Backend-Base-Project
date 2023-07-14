package service

import (
	"backend/internal/app/product/models"
	"backend/pkg/utils"
)

// ProductService is an interface that defines the contract for interacting with product-related functionality.
type ProductService interface {
	Create(request *models.CreateProductRequestDTO) (models.GetProductResponseDTO, error)
	Update(id string, request *models.UpdateProductRequestDTO) (models.GetProductResponseDTO, error)
	Delete(id string) (models.GetProductResponseDTO, error)
	GetPagination(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
	GetProductByID(id string) (models.GetProductResponseDTO, error)
	// Add other service methods here
}
