package service

import (
	"backend/internal/app/transaction/models"
	"backend/pkg/utils"
)

// TransactionService is an interface that defines the contract for interacting with transaction-related functionality.
type TransactionService interface {
	CreateTransaction(transaction *models.CreateTransactionDTO) (*models.GetTransactionDTO, error)
	GetTransactionByID(id string) (*models.GetTransactionDTO, error)
	GetPagination(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
	// Add other products methods here
}
