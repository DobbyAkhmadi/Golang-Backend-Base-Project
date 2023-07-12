package repository

import (
	"backend/internal/app/product/models"
	"backend/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ProductRepository is an interface that defines the contract for accessing and manipulating Product data.
type ProductRepository interface {
	// Save basic crud
	Save(model *models.Product) (*models.Product, error)
	Update(model *models.Product) (*models.Product, error)
	GetByID(id string) (*models.Product, error)
	Delete(id string) error
	GetPagination(paginate utils.SetPaginationDto) ([]*models.Product, int64, error)

	// BulkDelete custom query include array data
	BulkDelete(id []string) error
	GetByColumns(columns []string, values []string) ([]*models.Product, error)
}

// dbProductRepository implements the ProductRepository interface.
type dbProductRepository struct {
	connection *gorm.DB
}

func (db dbProductRepository) Delete(id string) error {
	product := &models.Product{}
	if err := db.connection.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func (db dbProductRepository) BulkDelete(id []string) error {
	if err := db.connection.Where("id IN (?)", id).Delete(&models.Product{}).Error; err != nil {
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
	var users []*models.Product
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
	if err := query.Offset(offset).Limit(paginate.PageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (db dbProductRepository) GetByColumns(columns []string, values []string) ([]*models.Product, error) {
	var products []*models.Product
	query := db.connection

	// Build the query dynamically based on the specified columns and values
	for i := 0; i < len(columns); i++ {
		query = query.Where(columns[i]+" = ?", values[i])
	}

	// Execute the query and retrieve the matching products
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// NewProductRepository creates a new instance of the ProductRepository.
// It performs necessary database migrations and generates fake data if in the development environment.
func NewProductRepository() ProductRepository {
	// Check if the database connection is already established
	if utils.DB == nil {
		// Connect to the database
		database, _ := utils.Connect()
		if database != nil {
			log.Error(database)
		}
	}

	// Perform auto-migration for Product table
	model := models.Product{}
	err := model.AutoMigrate(utils.DB)
	if err != nil {
		panic(err)
	}
	log.Info("Migration completed successfully!")

	// Return the dbProductRepository instance
	return &dbProductRepository{
		connection: utils.DB,
	}
}

// Save saves a new service.product into the database.
func (db dbProductRepository) Save(model *models.Product) (*models.Product, error) {
	if err := db.connection.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (db dbProductRepository) Update(model *models.Product) (*models.Product, error) {
	if err := db.connection.Save(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
