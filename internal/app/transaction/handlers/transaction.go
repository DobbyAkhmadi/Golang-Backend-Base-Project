package handlers

import (
	"backend/internal/app/transaction/models"
	"backend/internal/app/transaction/service"
	"backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	// Parse and validate the request body
	transaction := new(models.CreateTransactionDTO)

	if err := c.BodyParser(transaction); err != nil {
		// Return a response with a parsing error
		log.Info(transaction)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Create the service.transaction using the service.transaction service
	data, err := h.transactionService.CreateTransaction(transaction)

	if err != nil {
		// Return a response with an error message
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return c.JSON(utils.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

func (h *TransactionHandler) GetPaginationTransaction(ctx *fiber.Ctx) error {
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
	response, err := h.transactionService.GetPagination(pagination)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(response)

}

func (h *TransactionHandler) GetTransactionById(ctx *fiber.Ctx) error {
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
	data, err := h.transactionService.GetTransactionByID(id)
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
		Data:   data,
	})
}
