package models

type CreateProductRequestDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int64  `json:"stock"`
}

type UpdateProductRequestDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int64  `json:"stock"`
}

type CreateProductResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int64  `json:"stock"`
}

type GetProductResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int64  `json:"stock"`
}
