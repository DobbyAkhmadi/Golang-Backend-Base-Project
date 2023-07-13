package service

import "backend/internal/app/user/repository"

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
