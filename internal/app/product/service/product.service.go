package service

import (
	"backend/internal/app/product/models"
	"backend/pkg/utils"
)

type ProductService interface {
	Create(request *models.CreateProductRequestDTO) (models.GetProductResponseDTO, error)
	Update(request *models.UpdateProductRequestDTO) (models.GetProductResponseDTO, error)
	Delete(id string) (models.GetProductResponseDTO, error)
	GetPagination(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
	GetProductByID(id string) (models.GetProductResponseDTO, error)
	// Add other service methods here
}
