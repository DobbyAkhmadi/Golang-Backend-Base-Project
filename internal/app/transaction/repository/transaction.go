package repository

import (
	models2 "backend/internal/app/transaction/models"
	"backend/pkg/utils"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models2.Transaction, details []*models2.TransactionDetail) (uuid.UUID, error)
	GetTransactionByID(id uuid.UUID) (models2.Transaction, []*models2.TransactionDetail, error)
	GetPagination(paginate utils.SetPaginationDto) ([]*models2.Transaction, []*models2.TransactionDetail, int64, error)
}

type dbTransactionRepository struct {
	connection *gorm.DB
}

func (r *dbTransactionRepository) GetPagination(paginate utils.SetPaginationDto) ([]*models2.Transaction, []*models2.TransactionDetail, int64, error) {
	offset := (paginate.PageIndex - 1) * paginate.PageSize
	var transaction []*models2.Transaction
	var details []*models2.TransactionDetail
	var total int64

	query := r.connection.Model(&models2.Transaction{}).Find(&transaction)
	if query.Error != nil {
		return transaction, details, total, query.Error
	}
	// Retrieve the details associated with the transactions
	query = r.connection.Model(&models2.TransactionDetail{}).Where("transaction_id IN (?)", transaction).Find(&details)
	if query.Error != nil {
		return transaction, details, total, query.Error
	}
	// Apply sorting
	if paginate.SortBy != "" {
		query = query.Order(paginate.SortBy)
	}

	// Apply global search
	if paginate.GlobalSearch != "" {
		search := "%" + paginate.GlobalSearch + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", search, search)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, 0, err
	}

	// Apply pagination
	if err := query.Offset(offset).Limit(paginate.PageSize).Find(&transaction).Error; err != nil {
		return nil, nil, 0, err
	}

	return transaction, details, total, nil
}

func (r *dbTransactionRepository) GetTransactionByID(id uuid.UUID) (models2.Transaction, []*models2.TransactionDetail, error) {
	var transaction models2.Transaction
	var details []*models2.TransactionDetail

	// Retrieve the transaction by ID from the database
	if err := r.connection.First(&transaction, id).Error; err != nil {
		return transaction, details, err
	}

	// Retrieve the transaction details associated with the transaction
	if err := r.connection.Where("transaction_id = ?", id).Find(&details).Error; err != nil {
		return transaction, details, err
	}

	return transaction, details, nil
}

func NewTransactionRepository() TransactionRepository {
	// Check if the database connection is already established
	if utils.DB == nil {
		// Connect to the database
		database, _ := utils.Connect()
		if database != nil {
			log.Error(database)
		}
	}

	// Perform auto-migration for Transaction table
	transaction := models2.Transaction{}
	err := transaction.AutoMigrate(utils.DB)
	if err != nil {
		panic(err)
	}

	// Perform auto-migration for TransactionDetail table
	transactionDetail := models2.TransactionDetail{}
	err_ := transactionDetail.AutoMigrate(utils.DB)
	if err_ != nil {
		panic(err_)
	}

	log.Info("Migration completed successfully!")

	// Return the dbProductRepository instance
	return &dbTransactionRepository{
		connection: utils.DB,
	}
}
func (r *dbTransactionRepository) CreateTransaction(transaction *models2.Transaction, details []*models2.TransactionDetail) (uuid.UUID, error) {
	err := r.connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		for _, detail := range details {
			detail.TransactionID = transaction.ID
			if err := tx.Create(detail).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return uuid.Nil, nil
	}

	return transaction.ID, nil
}
