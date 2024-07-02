package service

import (
	"backend/internal/app/product/models"
	"backend/pkg/utils"
)

type IProductService interface {
	Create(request *models.CreateProductRequestDTO) (models.GetProductResponseDTO, error)
	Update(id string, request *models.UpdateProductRequestDTO) (models.GetProductResponseDTO, error)
	Delete(id string) (models.GetProductResponseDTO, error)
	Restore(id string) (models.GetProductResponseDTO, error)
	GetPagination(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
	GetProductByID(id string) (models.GetProductResponseDTO, error)
}
