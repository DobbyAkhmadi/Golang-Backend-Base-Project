package service

import (
	"backend/internal/app/user/models"
	"backend/internal/app/user/repository"
)

// UserAuthServiceImpl represents an implementation of the UserService interface.
type UserAuthServiceImpl struct {
	userRepository repository.UserRepository
}

func (u UserAuthServiceImpl) Login(dto models.AuthLoginRequestDTO) (models.AuthLoginResponseDTO, error) {
	//TODO implement me
	panic("implement me")
}

// NewUserAuthService creates a new instance of UserServiceImpl.
func NewUserAuthService(userRepository *repository.UserRepository) *UserAuthServiceImpl {
	return (*UserAuthServiceImpl)(&UserServiceImpl{
		userRepository: *userRepository,
	})
}

func (u *UserServiceImpl) Login(dto models.AuthLoginRequestDTO) (models.AuthLoginResponseDTO, error) {

	return models.AuthLoginResponseDTO{}, nil
}
