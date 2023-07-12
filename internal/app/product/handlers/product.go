package handlers

import (
	"backend/internal/app/product/models"
	"backend/internal/app/product/service"
	"backend/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateNewProduct creates a new service.product.
// @Summary Create a new service.product
// @Description Create a new service.product with the provided request data
// @Tags companies
// @Accept json
// @Produce json
// @Param request body models.CreateProductRequestDTO true "Request body containing service.product details"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/v1/service.product [post]
func (h *ProductHandler) CreateNewProduct(ctx *fiber.Ctx) error {
	// request new data from http
	request := new(models.CreateProductRequestDTO)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Errors: err,
		})
	}
	// Validate the Product struct
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Errors: validationErrors,
		})
	}

	// Call the CreateProduct method of the productService
	response, err := h.productService.Create(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusConflict,
			Status: "cannot insert duplicate values",
			Errors: err,
		})
	}

	// return the response into JSON
	return ctx.JSON(utils.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (h *ProductHandler) UpdateExistingProduct(c *fiber.Ctx) error {
	// request new data from http
	request := new(models.UpdateProductResponseDTO)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Errors: err,
		})
	}
	// Validate the Product struct
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Errors: validationErrors,
		})
	}

	// Call the UpdateProduct method of the productService
	response, err := h.productService.Update(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusNotFound,
			Status: "service.product not found",
			Errors: err,
		})
	}

	// return the response into JSON
	return c.JSON(utils.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

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
			"error": "Missing required query parameters",
		})
	}

	// Convert the query parameters to their respective types
	pageIndexInt, err := strconv.Atoi(pageIndex)
	if err != nil || pageIndexInt < 0 {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page index",
		})
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt <= 0 {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page size",
		})
	}

	// Create the pagination DTO with the request parameters
	pagination := utils.SetPaginationDto{
		PageIndex:    pageIndexInt,
		PageSize:     pageSizeInt,
		GlobalSearch: globalSearch,
		SortBy:       sortBy,
	}

	// Retrieve the paginated products from the service
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

func (h *ProductHandler) GetProductByID(ctx *fiber.Ctx) error {
	// Get the service.product ID from the request parameters
	id := ctx.Params("id")

	// Validate the service.product ID
	if id == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid service.product ID",
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
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   product,
	})
}

func (h *ProductHandler) DeleteExistingProduct(ctx *fiber.Ctx) error {
	// Get the service.product ID from the request parameters
	id := ctx.Params("id")

	// Validate the service.product ID
	if id == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid service.product ID",
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
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   product,
	})
}