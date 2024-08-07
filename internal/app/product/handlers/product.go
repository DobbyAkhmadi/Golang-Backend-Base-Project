package handlers

import (
	"backend/internal/app/product/models"
	srv "backend/internal/app/product/service"
	"backend/pkg/utils"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ProductHandler struct {
	productService srv.IProductService
}

func NewProductHandler(productService srv.IProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateNewProduct creates a new product.
// @Summary Create a new product
// @Description Create a new product with the provided request data
// @Tags Product
// @Accept json
// @Produce json
// @Param request body models.CreateProductRequestDTO true "Request body containing product details"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/product [post]
func (h *ProductHandler) CreateNewProduct(ctx *fiber.Ctx) error {
	// request new data from http
	request := new(models.CreateProductRequestDTO)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: utils.StatusBadRequest,
			Errors: err,
		})
	}
	// Validate the Product struct
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: utils.StatusBadRequest,
			Errors: validationErrors,
		})
	}

	// Call the CreateProduct method of the productService
	response, err := h.productService.Create(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusConflict,
			Status: utils.StatusConflict,
			Errors: err,
		})
	}

	// return the response into JSON
	return ctx.JSON(utils.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusOK,
		Payload: response,
	})
}

// UpdateExistingProduct updates existing product.
// @Summary Update existing product
// @Description Update existing product with the provided request data
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body models.UpdateProductRequestDTO true "Request body containing product details"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/product [put]
func (h *ProductHandler) UpdateExistingProduct(ctx *fiber.Ctx) error {
	// Get the service.product ID from the request parameters
	id := ctx.Params("id")

	// Validate the service.product ID
	if id == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.StatusBadRequest,
		})
	}
	// request new data from http
	request := new(models.UpdateProductRequestDTO)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: utils.StatusBadRequest,
			Errors: err,
		})
	}
	// Validate the Product struct
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: utils.StatusBadRequest,
			Errors: validationErrors,
		})
	}

	// Call the UpdateProduct method of the productService
	response, err := h.productService.Update(id, request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusNotFound,
			Status: utils.StatusNotFound,
			Errors: err,
		})
	}

	// return the response into JSON
	return ctx.JSON(utils.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusOK,
		Payload: response,
	})
}

// GetPaginationProduct get pagination all product.
// @Summary Get Pagination Product
// @Description Get All Paginated product with the provided request data
// @Tags Product
// @Accept json
// @Produce json
// @Param page_index query int false "Page index" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param search query string false "Global search term"
// @Param sort_by query string false "Sort by field"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/product [get]
func (h *ProductHandler) GetPaginationProduct(ctx *fiber.Ctx) error {
	// Parse and validate the request query parameters
	pageIndex := ctx.Query("page_index")
	pageSize := ctx.Query("page_size")
	globalSearch := ctx.FormValue("search")
	sortBy := ctx.FormValue("sort_by")

	// Perform any necessary validation on the query parameters
	if pageIndex == "" || pageSize == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.MissingRequiredParams,
		})
	}

	// Convert the query parameters to their respective types
	pageIndexInt, err := strconv.Atoi(pageIndex)
	if err != nil || pageIndexInt < 0 {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.InvalidPageIndex,
		})
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt <= 0 {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.InvalidPageSize,
		})
	}

	// Create the pagination DTO with the request parameters
	pagination := utils.SetPaginationDto{
		PageIndex:    pageIndexInt,
		PageSize:     pageSizeInt,
		GlobalSearch: globalSearch,
		SortBy:       sortBy,
	}

	// Retrieve the paginated product from the service
	response, err := h.productService.GetPagination(pagination)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(response)

}

// GetProductByID get a product by ID.
// @Summary Get Product by ID
// @Description Get a product by the provided ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/product/{id} [get]
func (h *ProductHandler) GetProductByID(ctx *fiber.Ctx) error {
	// Get the service.product ID from the request parameters
	id := ctx.Params("id")

	// Validate the service.product ID
	if id == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.StatusBadRequest,
		})
	}

	// Retrieve the service.product by ID from the service
	product, err := h.productService.GetProductByID(id)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(utils.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusOK,
		Payload: product,
	})
}

// DeleteProductByID delete a product by ID.
// @Summary Delete Product by ID
// @Description Delete a product by the provided ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/product/{id} [delete]
func (h *ProductHandler) DeleteProductByID(ctx *fiber.Ctx) error {
	// Get the service.product ID from the request parameters
	id := ctx.Params("id")

	// Validate the service.product ID
	if id == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.StatusBadRequest,
		})
	}

	// Delete the service.product by ID using the service
	product, err := h.productService.Delete(id)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(utils.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusOK,
		Payload: product,
	})
}

// RestoreProductByID delete a product by ID.
// @Summary Restore Product by ID
// @Description Restore a product by the provided ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/product/{id} [patch]
func (h *ProductHandler) RestoreProductByID(ctx *fiber.Ctx) error {
	// Get the service.product ID from the request parameters
	id := ctx.Params("id")

	// Validate the service.product ID
	if id == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.StatusBadRequest,
		})
	}

	// Restore the service.product by ID using the service
	product, err := h.productService.Restore(id)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(utils.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusOK,
		Payload: product,
	})
}
