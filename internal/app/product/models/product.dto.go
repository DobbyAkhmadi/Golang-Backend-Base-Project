package models

// CreateProductRequestDTO represents the request structure for creating a product.
type CreateProductRequestDTO struct {
	Name        string `json:"name" validate:"required"`        // Name of the product (required)
	Description string `json:"description" validate:"required"` // Description of the product (required)
	Stock       int64  `json:"stock" validate:"required,min=0"` // Stock count of the product (required, minimum value: 0)
}

// UpdateProductRequestDTO represents the request structure for updating a product.
type UpdateProductRequestDTO struct {
	ID          string `json:"id" validate:"required"`          // ID of the product to update (required)
	Name        string `json:"name" validate:"required"`        // Name of the product (required)
	Description string `json:"description" validate:"required"` // Description of the product (required)
	Stock       int64  `json:"stock" validate:"required,min=0"` // Stock count of the product (required, minimum value: 0)
}

// GetProductResponseDTO represents the response structure for a product.
type GetProductResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int64  `json:"stock"`
}
