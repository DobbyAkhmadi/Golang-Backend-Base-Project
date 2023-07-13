package models

// AuthLoginRequestDTO represents the request structure for user login.
type AuthLoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"` // Email of the user (required, must be a valid email address)
	Password string `json:"password" validate:"required"`    // Password of the user (required)
}

// AuthLoginResponseDTO represents the response structure for user login.
type AuthLoginResponseDTO struct {
	Username string           `json:"username"` // Username of the user
	Email    string           `json:"email" `   // Email of the user
	Token    string           `json:"token"`    // Token of the user and auto generated when login
	Roles    RolesResponseDto `json:"roles"`    // Roles of the user get authority from current user
}

// RolesResponseDto represents the response structure for user roles.
type RolesResponseDto struct {
	ID       int    `json:"id"`         // ID of the roles
	RoleName string `json:"role_name" ` // RoleName of the roles
}
