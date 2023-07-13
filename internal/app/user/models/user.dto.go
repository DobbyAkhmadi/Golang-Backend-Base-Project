package models

// CreateUserRequestDTO represents the request structure for new user
type CreateUserRequestDTO struct {
	FirstName   string `json:"first_name,omitempty"`             // FirstName of the user (optional)
	LastName    string `json:"last_name,omitempty"`              // LastName of the user (optional)
	PhoneNumber string `json:"phone,omitempty"`                  // PhoneNumber of the user (optional)
	Address     string `json:"address,omitempty"`                // Address of the user (optional)
	Username    string `json:"username,omitempty"`               // Username of the user (optional)
	Password    string `json:"password" validate:"required"`     // Password of the user (required)
	Email       string `json:"email,omitempty" validate:"email"` // Email of the user (optional, must be a valid email address)
}

// UpdateUserRequestDTO represents the request structure for existing user
type UpdateUserRequestDTO struct {
	FirstName   string `json:"first_name,omitempty"`             // FirstName of the user (optional)
	LastName    string `json:"last_name,omitempty"`              // LastName of the user (optional)
	PhoneNumber string `json:"phone,omitempty"`                  // PhoneNumber of the user (optional)
	Address     string `json:"address,omitempty"`                // Address of the user (optional)
	Username    string `json:"username,omitempty"`               // Username of the user (optional)
	Password    string `json:"password" validate:"required"`     // Password of the user (required)
	Email       string `json:"email,omitempty" validate:"email"` // Email of the user (optional, must be a valid email address)
}

// UpdateUserRequestPasswordDTO represents the request structure for existing user
type UpdateUserRequestPasswordDTO struct {
	Email    string `json:"email,omitempty" validate:"email"` // Email of the user (optional, must be a valid email address)
	Password string `json:"password" validate:"required"`     // Password of the user (required)
}

// GetUserResponseDTO represents the response structure for new/existing user
type GetUserResponseDTO struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`                       // FirstName of the user
	LastName    string `json:"last_name"`                        // LastName of the user
	PhoneNumber string `json:"phone"`                            // PhoneNumber of the user
	Address     string `json:"address"`                          // Address of the user
	Username    string `json:"username"`                         // Username of the user
	Email       string `json:"email,omitempty" validate:"email"` // Email of the user
}
