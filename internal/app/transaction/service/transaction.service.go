package service

import (
	"backend/internal/app/transaction/models"
	"backend/pkg/utils"
)

type TransactionService interface {
	CreateTransaction(transaction *models.CreateTransactionDTO) (*models.GetTransactionDTO, error)
	GetTransactionByID(id string) (*models.GetTransactionDTO, error)
	GetPagination(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
	// Add other service methods here
}
