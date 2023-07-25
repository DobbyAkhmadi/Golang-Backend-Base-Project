package handlers

import (
	"backend/internal/app/region/service"
	"backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type RegionHandler struct {
	regionService service.RegionService
}

func NewRegionHandler(regionService service.RegionService) *RegionHandler {
	return &RegionHandler{
		regionService: regionService,
	}
}

// GetPaginationVillage get pagination all village.
// @Summary Get Pagination village
// @Description Get All Paginated village with the provided request data
// @Tags Village
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
// @Router /api/v1/village [get]
func (h *RegionHandler) GetPaginationVillage(ctx *fiber.Ctx) error {
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

	// Retrieve the paginated village from the service
	response, err := h.regionService.GetPaginationVillage(pagination)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(response)

}

// GetPaginationProvince get pagination all province.
// @Summary Get Pagination province
// @Description Get All Paginated village with the provided request data
// @Tags Province
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
// @Router /api/v1/province [get]
func (h *RegionHandler) GetPaginationProvince(ctx *fiber.Ctx) error {
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

	// Retrieve the paginated province from the service
	response, err := h.regionService.GetPaginationProvince(pagination)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(response)

}

// GetPaginationDistrict get pagination all district.
// @Summary Get Pagination district
// @Description Get All Paginated village with the provided request data
// @Tags District
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
// @Router /api/v1/district [get]
func (h *RegionHandler) GetPaginationDistrict(ctx *fiber.Ctx) error {
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

	// Retrieve the paginated district from the service
	response, err := h.regionService.GetPaginationDistrict(pagination)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(response)

}

// GetPaginationRegency get pagination all regency.
// @Summary Get Pagination regency
// @Description Get All Paginated village with the provided request data
// @Tags Regency
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
// @Router /api/v1/regency [get]
func (h *RegionHandler) GetPaginationRegency(ctx *fiber.Ctx) error {
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

	// Retrieve the paginated regency from the service
	response, err := h.regionService.GetPaginationRegency(pagination)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(response)

}
