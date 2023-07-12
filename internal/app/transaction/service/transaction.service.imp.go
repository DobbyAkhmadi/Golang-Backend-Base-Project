package service

import (
	"backend/internal/app/transaction/models"
	"backend/internal/app/transaction/repository"
	"backend/pkg/utils"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(productRepository *repository.TransactionRepository) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		transactionRepository: *productRepository,
	}
}

func (s *TransactionServiceImpl) GetPagination(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error) {
	products, _, total, err := s.transactionRepository.GetPagination(paginate)
	if err != nil {
		return utils.GetGlobalResponsePaginationDto{}, err
	}

	generate := utils.GetGlobalResponsePaginationDto{
		Code:      200,
		Status:    "OK",
		Data:      products,
		PageIndex: paginate.PageIndex,
		PageSize:  paginate.PageSize,
		TotalRows: total,
	}

	return generate, nil
}

func (s *TransactionServiceImpl) GetTransactionByID(id string) (*models.GetTransactionDTO, error) {
	repo, data, _ := s.transactionRepository.GetTransactionByID(uuid.MustParse(id))

	response := new(models.GetTransactionDTO)
	//	response.Code = fiber.StatusOK
	//response.Status = "OK"
	response.ID = repo.ID.String()
	response.TransactionDate = repo.TransactionDate
	response.Data = convertToDTO(data)

	return response, nil
}

func (s *TransactionServiceImpl) CreateTransaction(transaction *models.CreateTransactionDTO) (*models.GetTransactionDTO, error) {
	// Perform any additional validation or business logic before creating the service.transaction

	head := models.Transaction{
		TransactionDate: time.Now(),
	}
	var transactionDetails []*models.TransactionDetail
	for _, detail := range transaction.Data {
		row := &models.TransactionDetail{
			ProductID: uuid.MustParse(detail.ProductID),
			Qty:       detail.Qty,
		}
		transactionDetails = append(transactionDetails, row)
	}

	id, err := s.transactionRepository.CreateTransaction(&head, transactionDetails)
	if err != nil {
		return nil, err
	}

	repo, data, _ := s.transactionRepository.GetTransactionByID(id)
	log.Info(repo)
	response := new(models.GetTransactionDTO)
	response.ID = repo.ID.String()
	response.TransactionDate = repo.TransactionDate
	response.Data = convertToDTO(data)

	return response, nil
}

func convertToDTO(model []*models.TransactionDetail) []*models.GetTransactionDetailDto {
	var dto []*models.GetTransactionDetailDto

	for _, m := range model {
		d := &models.GetTransactionDetailDto{
			ID:            m.ID.String(),
			TransactionID: m.TransactionID.String(),
			ProductID:     m.ProductID.String(),
			Qty:           m.Qty,
		}
		dto = append(dto, d)
	}

	return dto
}
