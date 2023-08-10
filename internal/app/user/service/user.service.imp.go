package service

import (
	models2 "backend/internal/app/user/models"
	"backend/internal/app/user/repository"
	"backend/pkg/utils"
	"github.com/google/uuid"
)

// UserServiceImpl represents an implementation of the UserService interface.
type UserServiceImpl struct {
	userRepository repository.UserRepository
}

// NewUserService  creates a new instance of UserServiceImpl.
func NewUserService(userRepository *repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: *userRepository,
	}
}

func (u *UserServiceImpl) GetPagination(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error) {
	// Retrieve paginated users from the repository
	users, total, err := u.userRepository.GetPagination(paginate)
	if err != nil {
		return utils.GetGlobalResponsePaginationDto{}, err
	}
	generate := utils.GetGlobalResponsePaginationDto{Header: utils.HeaderDto{
		Milliseconds: utils.GetCurrentLatency(),
		Message:      "Request Successfully",
	},
		Code:      200,
		Status:    "OK",
		Data:      users,
		PageIndex: paginate.PageIndex,
		PageSize:  paginate.PageSize,
		TotalRows: total,
	}

	return generate, nil
}

func (u *UserServiceImpl) Delete(id string) (models2.GetUserResponseDTO, error) {
	// Check if the service.user exists
	existingUser, err := u.userRepository.GetByID(id)
	if err != nil {
		return models2.GetUserResponseDTO{}, err
	}

	// Delete the service.users
	err = u.userRepository.Delete(id)
	if err != nil {
		return models2.GetUserResponseDTO{}, err
	}

	// Convert the deleted service.users to the response DTO
	dto := convertToDTO(existingUser)

	return dto, nil
}

func (u *UserServiceImpl) Create(request *models2.CreateUserRequestDTO) (models2.GetUserResponseDTO, error) {
	// Add business logic for creating a User
	// Validate the User, perform any necessary transformations, and interact with the repository
	// Return any relevant errors

	newComp := new(models2.User)
	newComp.ID = uuid.New()
	newComp.FirstName = request.FirstName
	newComp.LastName = request.LastName
	newComp.PhoneNumber = request.PhoneNumber
	newComp.Address = request.Address
	newComp.Username = request.Username
	newComp.Email = request.Email
	newComp.Password = request.Password

	// return error when insert duplicate values
	result, err := u.userRepository.Save(newComp)
	if err != nil {
		return models2.GetUserResponseDTO{}, err
	}
	// Convert the model to the response DTO
	dto := convertToDTO(result)

	return dto, nil
}

func (u *UserServiceImpl) Update(id string, request *models2.UpdateUserRequestDTO) (models2.GetUserResponseDTO, error) {
	// Add business logic for creating a User
	// Validate the User, perform any necessary transformations, and interact with the repository
	// Return any relevant errors

	// check existing data
	update, err := u.userRepository.GetByID(id)
	if err != nil {
		return models2.GetUserResponseDTO{}, err
	}
	update.FirstName = request.FirstName
	update.LastName = request.LastName
	update.PhoneNumber = request.PhoneNumber
	update.Address = request.Address
	update.Username = request.Username
	update.Email = request.Email
	update.Password = request.Password

	// return error when insert duplicate values
	result, err := u.userRepository.Update(update)
	if err != nil {
		return models2.GetUserResponseDTO{}, err
	}

	// Convert the model to the response DTO
	dto := convertToDTO(result)

	return dto, nil
}
func (u *UserServiceImpl) GetUserByID(id string) (models2.GetUserResponseDTO, error) {
	// Retrieve the service.user by ID from the repository
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return models2.GetUserResponseDTO{}, err
	}

	// Convert the service.user to the response DTO format
	dto := convertToDTO(user)

	return dto, nil
}

func convertToDTO(model *models2.User) models2.GetUserResponseDTO {
	dto := models2.GetUserResponseDTO{
		ID:          model.ID.String(),
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		PhoneNumber: model.PhoneNumber,
		Address:     model.Address,
		Username:    model.Username,
		Email:       model.Email,
	}

	return dto
}
