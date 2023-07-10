package repository

import (
	"backend/internal/app/company/models"
	"backend/pkg/utils"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"strings"
)

// CompanyRepository is an interface that defines the contract for accessing and manipulating Company data.
type CompanyRepository interface {
	Save(company models.Company) (uuid.UUID, error)
	Update(models.Company) error
	Delete(models.Company) error
	FindAll() []*models.Company
	FindByID(companyID uuid.UUID) (*models.Company, error)
	DeleteByID(companyID uuid.UUID) error
	FindByName(name string) (*models.Company, error)
	FindByField(fieldName, fieldValue string) (*models.Company, error)
	UpdateSingleField(company models.Company, fieldName, fieldValue string) error
}

// companyDatabase implements the CompanyRepository interface.
type companyDatabase struct {
	connection *gorm.DB
}

// NewCompanyRepository creates a new instance of the CompanyRepository.
// It performs necessary database migrations and generates fake data if in the development environment.
func NewCompanyRepository() CompanyRepository {
	// Check if the database connection is already established
	if utils.DB == nil {
		// Connect to the database
		database, _ := utils.Connect()
		if database != nil {
			log.Error(database)
		}

		// Perform auto-migration for Company table
		model := models.Company{}
		err := model.AutoMigrate(utils.DB)
		if err != nil {
			panic(err)
		}

		// Perform auto-migration for CompanyCategory table
		model_ := models.CompanyCategory{}
		err_ := model_.AutoMigrate(utils.DB)
		if err_ != nil {
			panic(err_)
		}

		// Generate fake data when in the development environment
		configEnv := strings.ToLower(os.Getenv("SERVER_MODE"))
		if configEnv == "dev" {
			err = generateFakeData(database)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("Migration completed successfully!")
	}

	// Return the companyDatabase instance
	return &companyDatabase{
		connection: utils.DB,
	}
}

func generateFakeData(db *gorm.DB) error {
	gofakeit.Seed(0) // Set the random seed for consistent results

	// Generate fake company category
	category := models.CompanyCategory{
		Name: gofakeit.RandomString([]string{"Category A", "Category B", "Category C"}),
	}

	if err := db.Create(&category).Error; err != nil {
		return err
	}
	// Generate fake companies
	for i := 0; i < 10; i++ {
		company := models.Company{
			Name:              gofakeit.Company(),
			Description:       gofakeit.Sentence(10),
			Website:           gofakeit.URL(),
			CompanyCategoryID: category.ID,
		}
		if err := db.Create(&company).Error; err != nil {
			return err
		}

	}

	return nil
}

// DeleteByID deletes a company by its ID.
func (db companyDatabase) DeleteByID(companyID uuid.UUID) error {
	company := models.Company{}
	company.ID = companyID
	result := db.connection.Delete(&company)
	return result.Error
}

// Save saves a new company into the database.
func (db companyDatabase) Save(company models.Company) (uuid.UUID, error) {
	result := db.connection.Create(&company)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return company.ID, nil
}

// Update updates an existing company in the database.
func (db companyDatabase) Update(company models.Company) error {
	result := db.connection.Save(&company)
	return result.Error
}

// Delete deletes a company from the database.
func (db companyDatabase) Delete(company models.Company) error {
	result := db.connection.Delete(&company)
	return result.Error
}

// FindAll retrieves all companies from the database.
func (db companyDatabase) FindAll() []*models.Company {
	var companies []*models.Company
	db.connection.Preload(clause.Associations).Find(&companies)
	return companies
}

// FindByID retrieves a company by its ID from the database.
func (db companyDatabase) FindByID(companyID uuid.UUID) (*models.Company, error) {
	var company models.Company
	result := db.connection.Preload(clause.Associations).Find(&company, "id = ?", companyID)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &company, nil
	}
	return nil, nil
}

// FindByName retrieves a company by its name from the database.
func (db companyDatabase) FindByName(name string) (*models.Company, error) {
	var company models.Company
	result := db.connection.Preload(clause.Associations).Find(&company, "name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &company, nil
	}
	return nil, nil
}

// FindByField retrieves a company by a specific field value from the database.
func (db companyDatabase) FindByField(fieldName, fieldValue string) (*models.Company, error) {
	var company models.Company
	result := db.connection.Preload(clause.Associations).Find(&company, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &company, nil
	}
	return nil, nil
}

// UpdateSingleField updates a single field of a company in the database.
func (db companyDatabase) UpdateSingleField(company models.Company, fieldName, fieldValue string) error {
	result := db.connection.Model(&company).Update(fieldName, fieldValue)
	return result.Error
}
