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

type CompanyCategoryRepository interface {
	Save(companyCategory models.CompanyCategory) (uuid.UUID, error)
	Update(models.CompanyCategory) error
	Delete(models.CompanyCategory) error
	FindAll() []*models.CompanyCategory
	FindByID(companyCategoryID uuid.UUID) (*models.CompanyCategory, error)
	DeleteByID(companyCategoryID uuid.UUID) error
	FindByName(name string) (*models.CompanyCategory, error)
	FindByField(fieldName, fieldValue string) (*models.CompanyCategory, error)
	UpdateSingleField(companyCategory models.CompanyCategory, fieldName, fieldValue string) error
}
type companyCategoryDatabase struct {
	connection *gorm.DB
}

func NewCompanyCategoryRepository() CompanyCategoryRepository {
	if entity.DB == nil {
		connect, _ := entity.Connect()
		if connect != nil {
			log.Error(connect)
		}
		model := models.CompanyCategory{}
		err := model.AutoMigrate(entity.DB)
		if err != nil {
			panic(err)
		}
	}
	return &companyCategoryDatabase{
		connection: entity.DB,
	}
}

func (db companyCategoryDatabase) DeleteByID(companyCategoryID uuid.UUID) error {
	companyCategory := models.CompanyCategory{}
	companyCategory.ID = companyCategoryID
	result := db.connection.Delete(&companyCategory)
	return result.Error
}

func (db companyCategoryDatabase) Save(companyCategory models.CompanyCategory) (uuid.UUID, error) {
	result := db.connection.Create(&companyCategory)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return companyCategory.ID, nil
}

func (db companyCategoryDatabase) Update(companyCategory models.CompanyCategory) error {
	result := db.connection.Save(&companyCategory)
	return result.Error
}

func (db companyCategoryDatabase) Delete(companyCategory models.CompanyCategory) error {
	result := db.connection.Delete(&companyCategory)
	return result.Error
}

func (db companyCategoryDatabase) FindAll() []*models.CompanyCategory {
	var companyCategories []*models.CompanyCategory
	db.connection.Preload(clause.Associations).Find(&companyCategories)
	return companyCategories
}

func (db companyCategoryDatabase) FindByID(companyCategoryID uuid.UUID) (*models.CompanyCategory, error) {
	var companyCategory models.CompanyCategory
	result := db.connection.Preload(clause.Associations).Find(&companyCategory, "id = ?", companyCategoryID)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &companyCategory, nil
	}
	return nil, nil
}

func (db companyCategoryDatabase) FindByName(name string) (*models.CompanyCategory, error) {
	var companyCategory models.CompanyCategory
	result := db.connection.Preload(clause.Associations).Find(&companyCategory, "name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &companyCategory, nil
	}
	return nil, nil
}

func (db companyCategoryDatabase) FindByField(fieldName, fieldValue string) (*models.CompanyCategory, error) {
	var companyCategory models.CompanyCategory
	result := db.connection.Preload(clause.Associations).Find(&companyCategory, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &companyCategory, nil
	}
	return nil, nil
}

func (db companyCategoryDatabase) UpdateSingleField(companyCategory models.CompanyCategory, fieldName, fieldValue string) error {
	result := db.connection.Model(&companyCategory).Update(fieldName, fieldValue)
	return result.Error
}
