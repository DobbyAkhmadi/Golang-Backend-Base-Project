package service

import (
	"backend/config"
	"backend/internal/app/user/models"
	"backend/internal/app/user/repository"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// UserAuthServiceImpl represents an implementation of the UserService interface.
type UserAuthServiceImpl struct {
	userRepository repository.UserRepository
}

// NewUserAuthService creates a new instance of UserServiceImpl.
func NewUserAuthService(userRepository *repository.UserRepository) *UserAuthServiceImpl {
	return (*UserAuthServiceImpl)(&UserServiceImpl{
		userRepository: *userRepository,
	})
}

func (u UserAuthServiceImpl) Login(dto *models.AuthLoginRequestDTO) (*models.AuthLoginResponseDTO, error) {
	login, err := u.userRepository.FindByEmail(dto.Email, dto.Password)

	if err != nil {
		return nil, err
	}

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = login.ID
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	secretConfig := config.Config.GetString("JWT.TOKEN")
	if secretConfig == "" {
		return nil, errors.New("missing JWT token configuration")
	}

	tokenString, err := tokenByte.SignedString([]byte(secretConfig))
	if err != nil {
		return nil, err
	}

	auth := new(models.AuthLoginResponseDTO)
	auth.Username = login.Username
	auth.Email = login.Email
	auth.Token = tokenString
	auth.Roles = models.RolesResponseDto{
		ID:       0,
		RoleName: "Admin",
	}

	// return the response into JSON
	return auth, nil
}
