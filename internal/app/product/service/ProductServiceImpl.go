package service

import (
	"backend/internal/app/product/models"
	repo "backend/internal/app/product/repository"
	"backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductServiceImpl struct {
	productRepository repo.IProductRepository
}

func NewProductService(productRepository *repo.IProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepository: *productRepository,
	}
}

func (s *ProductServiceImpl) GetPagination(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error) {
	// Retrieve paginated product from the repository
	products, total, err := s.productRepository.GetPagination(paginate)
	if err != nil {
		return utils.GetGlobalResponsePaginationDto{}, err
	}

	generate := utils.GetGlobalResponsePaginationDto{
		Header: utils.HeaderDto{
			Milliseconds: utils.GetCurrentLatency(),
			Message:      utils.StatusOK,
		},
		Code:      fiber.StatusOK,
		Status:    utils.StatusOK,
		Data:      products,
		PageIndex: paginate.PageIndex,
		PageSize:  paginate.PageSize,
		TotalRows: total,
	}

	return generate, nil
}

func (s *ProductServiceImpl) Restore(id string) (models.GetProductResponseDTO, error) {
	// Check if the service.product exists
	existingProduct, err := s.productRepository.GetByID(id)
	if err != nil {
		return models.GetProductResponseDTO{}, err
	}

	// Restore the service.product
	err = s.productRepository.Restore(id)
	if err != nil {
		return models.GetProductResponseDTO{}, err
	}

	// Convert the deleted service.product to the response DTO
	dto := convertToDTO(existingProduct)

	return dto, nil
}

func (s *ProductServiceImpl) Delete(id string) (models.GetProductResponseDTO, error) {
	// Check if the service.product exists
	existingProduct, err := s.productRepository.GetByID(id)
	if err != nil {
		return models.GetProductResponseDTO{}, err
	}

	// Delete the service.product
	err = s.productRepository.Delete(id)
	if err != nil {
		return models.GetProductResponseDTO{}, err
	}

	// Convert the deleted service.product to the response DTO
	dto := convertToDTO(existingProduct)

	return dto, nil
}

func (s *ProductServiceImpl) Create(request *models.CreateProductRequestDTO) (models.GetProductResponseDTO, error) {

	newComp := new(models.Product)
	newComp.ID = uuid.New()
	newComp.Name = request.Name
	newComp.Stock = request.Stock
	newComp.Description = request.Description

	// return error when insert duplicate values
	result, err := s.productRepository.Upsert(newComp)
	if err != nil {
		return models.GetProductResponseDTO{}, err
	}
	// Convert the model to the response DTO
	dto := convertToDTO(result)

	return dto, nil
}

func (s *ProductServiceImpl) Update(id string, request *models.UpdateProductRequestDTO) (models.GetProductResponseDTO, error) {

	updateProduct, err := s.productRepository.GetByID(id)
	if err != nil {
		return models.GetProductResponseDTO{}, err
	}
	updateProduct.Name = request.Name
	updateProduct.Description = request.Description
	updateProduct.Stock = request.Stock

	// return error when insert duplicate values
	result, err := s.productRepository.Upsert(updateProduct)
	if err != nil {
		return models.GetProductResponseDTO{}, err
	}

	// Convert the model to the response DTO
	dto := convertToDTO(result)

	return dto, nil
}
func (s *ProductServiceImpl) GetProductByID(id string) (models.GetProductResponseDTO, error) {
	// Retrieve the service.product by ID from the repository
	product, err := s.productRepository.GetByID(id)
	if err != nil {
		return models.GetProductResponseDTO{}, err
	}

	// Convert the service.product to the response DTO format
	dto := convertToDTO(product)

	return dto, nil
}

func convertToDTO(model *models.Product) models.GetProductResponseDTO {
	dto := models.GetProductResponseDTO{
		ID:          model.ID.String(),
		Name:        model.Name,
		Description: model.Description,
		Stock:       model.Stock,
	}

	return dto
}
