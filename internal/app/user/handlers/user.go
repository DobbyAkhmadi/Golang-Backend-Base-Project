package handlers

import (
	models2 "backend/internal/app/user/models"
	"backend/internal/app/user/service"
	"backend/pkg/utils"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// UserHandler represents a handler for user-related operations.
type UserHandler struct {
	userService     service.UserService
	userAuthService service.UserAuthService
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userService service.UserService, userAuthService service.UserAuthService) *UserHandler {
	return &UserHandler{
		userService:     userService,
		userAuthService: userAuthService,
	}
}

// LoginUser Login a new User.
// @Summary login a new User
// @Description login User with the provided request data
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.AuthLoginRequestDTO true "Request body containing User login"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/user/auth/login [post]
func (h *UserHandler) LoginUser(ctx *fiber.Ctx) error {
	// request new data from http
	request := new(models2.AuthLoginRequestDTO)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: utils.StatusBadRequest,
			Errors: err,
		})
	}
	// Validate the User struct
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

	// Call the CreateUser method of the userService
	response, err := h.userAuthService.Login(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: utils.StatusBadRequest,
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

// CreateNewUser creates a new User.
// @Summary Create a new User
// @Description Create a new User with the provided request data
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequestDTO true "Request body containing User details"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/user [post]
func (h *UserHandler) CreateNewUser(ctx *fiber.Ctx) error {
	// request new data from http
	request := new(models2.CreateUserRequestDTO)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: utils.StatusBadRequest,
			Errors: err,
		})
	}
	// Validate the User struct
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

	// Call the CreateUser method of the userService
	response, err := h.userService.Create(request)
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

// UpdateExistingUser updates existing User.
// @Summary Update existing User
// @Description Update existing User with the provided request data
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body models.UpdateUserRequestDTO true "Request body containing User details"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/User [put]
func (h *UserHandler) UpdateExistingUser(ctx *fiber.Ctx) error {
	// Get the products.user ID from the request parameters
	id := ctx.Params("id")

	// Validate the products.user ID
	if id == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid products.user ID",
		})
	}
	// request new data from http
	request := new(models2.UpdateUserRequestDTO)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: utils.StatusBadRequest,
			Errors: err,
		})
	}
	// Validate the user struct
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

	// Call the Update user method of the userService
	response, err := h.userService.Update(id, request)
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

// GetPaginationUser get pagination all user.
// @Summary Get Pagination user
// @Description Get All Paginated user with the provided request data
// @Tags User
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
// @Router /api/v1/user [get]
func (h *UserHandler) GetPaginationUser(ctx *fiber.Ctx) error {
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

	// Retrieve the paginated user from the products
	response, err := h.userService.GetPagination(pagination)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(response)

}

// GetUserByID get a user by ID.
// @Summary Get user by ID
// @Description Get a user by the provided ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/user/{id} [get]
func (h *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	// Get the products.user ID from the request parameters
	id := ctx.Params("id")

	// Validate the products.user ID
	if id == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid products.user ID",
		})
	}

	// Retrieve the products.user by ID from the products
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(utils.Response{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Payload: user,
	})
}

// DeleteUserByID delete a user by ID.
// @Summary Delete user by ID
// @Description Delete a user by the provided ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/user/{id} [delete]
func (h *UserHandler) DeleteUserByID(ctx *fiber.Ctx) error {
	// Get the products.user ID from the request parameters
	id := ctx.Params("id")

	// Validate the products.user ID
	if id == "" {
		// Return a response with a validation error
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid products.user ID",
		})
	}

	// Delete the products.user by ID using the products
	user, err := h.userService.Delete(id)
	if err != nil {
		// Return a response with an error message
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the response into JSON
	return ctx.JSON(utils.Response{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Payload: user,
	})
}
