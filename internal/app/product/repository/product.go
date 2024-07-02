package repository

import (
	"backend/internal/app/product/models"
	"backend/pkg/utils"
	"backend/platform/database"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// IProductRepository is an interface that defines the contract for accessing and manipulating Product data.
type IProductRepository interface {
	Upsert(model *models.Product) (*models.Product, error)
	GetByID(id string) (*models.Product, error)
	Delete(id string) error
	Restore(id string) error
	GetPagination(paginate utils.SetPaginationDto) ([]*models.Product, int64, error)
}

// dbProductRepository implements the ProductRepository interface.
type dbProductRepository struct {
	connection *gorm.DB
}

func (db dbProductRepository) Upsert(model *models.Product) (*models.Product, error) {
	if err := db.connection.Save(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (db dbProductRepository) Delete(id string) error {
	if err := db.connection.Where("id = ?", id).Update("is_active = 3", id).Error; err != nil {
		return err
	}
	return nil
}

func (db dbProductRepository) Restore(id string) error {
	if err := db.connection.Where("id = ?", id).Update("is_active = 1", id).Error; err != nil {
		return err
	}
	return nil
}

func (db dbProductRepository) GetByID(id string) (*models.Product, error) {
	product := &models.Product{}
	if err := db.connection.Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (db dbProductRepository) GetPagination(paginate utils.SetPaginationDto) ([]*models.Product, int64, error) {
	offset := (paginate.PageIndex - 1) * paginate.PageSize
	var myModel []*models.Product
	var total int64

	query := db.connection.Model(&models.Product{})

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
		return nil, 0, err
	}

	// Apply pagination
	if err := query.Offset(offset).Limit(paginate.PageSize).Find(&myModel).Error; err != nil {
		return nil, 0, err
	}

	return myModel, total, nil
}

// NewProductRepository creates a new instance of the ProductRepository.
// It performs necessary database migrations and generates fake data if in the development environment.
func NewProductRepository() IProductRepository {
	// Check if the database connection is already established
	if database.DB == nil {
		// Connect to the database
		_db, _ := database.Connect()
		if _db != nil {
			log.Error(_db)
		}
	}

	// Perform auto-migrations for Product table
	model := models.Product{}
	err := model.AutoMigrate(database.DB)
	if err != nil {
		panic(err)
	}
	log.Info("Migration completed successfully!")

	// Return the dbProductRepository instance
	return &dbProductRepository{
		connection: database.DB,
	}
}
