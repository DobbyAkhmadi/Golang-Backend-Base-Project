package repository

import (
	"backend/internal/app/company/models"
	"backend/pkg/entity"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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
type companyDatabase struct {
	connection *gorm.DB
}

func NewCompanyRepository() CompanyRepository {
	if entity.DB == nil {
		company, _ := entity.Connect()
		if company != nil {
			log.Error(company)
		}
		model := models.Company{}
		err := model.AutoMigrate(entity.DB)
		if err != nil {
			panic(err)
		}
	}
	return &companyDatabase{
		connection: entity.DB,
	}
}

func (db companyDatabase) DeleteByID(companyID uuid.UUID) error {
	company := models.Company{}
	company.ID = companyID
	result := db.connection.Delete(&company)
	return result.Error
}

func (db companyDatabase) Save(company models.Company) (uuid.UUID, error) {
	result := db.connection.Create(&company)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return company.ID, nil
}

func (db companyDatabase) Update(company models.Company) error {
	result := db.connection.Save(&company)
	return result.Error
}

func (db companyDatabase) Delete(company models.Company) error {
	result := db.connection.Delete(&company)
	return result.Error
}

func (db companyDatabase) FindAll() []*models.Company {
	var companies []*models.Company
	db.connection.Preload(clause.Associations).Find(&companies)
	return companies
}

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

func (db companyDatabase) UpdateSingleField(company models.Company, fieldName, fieldValue string) error {
	result := db.connection.Model(&company).Update(fieldName, fieldValue)
	return result.Error
}
