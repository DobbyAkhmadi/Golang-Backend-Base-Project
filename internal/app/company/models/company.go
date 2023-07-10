package models

import (
	"backend/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Company represents a company entity.
type Company struct {
	utils.UUIDPrimaryKey                 // Embeds the UUIDPrimaryKey struct for a UUID primary key field
	Name                 string          `json:"name" gorm:"uniqueIndex"`                                                    // Name of the company
	Description          string          `json:"description,omitempty"`                                                      // Description of the company (optional)
	Website              string          `json:"website,omitempty"`                                                          // Website URL of the company (optional)
	CompanyCategoryID    uuid.UUID       `gorm:"type:uuid"`                                                                  // Foreign key to CompanyCategory table using UUID type
	CompanyCategories    CompanyCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CompanyCategoryID"` // One-to-many relationship with CompanyCategory
	utils.ModelTimestamp                 // Embeds the ModelTimestamp struct for timestamp fields
}

// CompanyCategory represents a category of a company.
type CompanyCategory struct {
	utils.UUIDPrimaryKey        // Embeds the UUIDPrimaryKey struct for a UUID primary key field
	Name                 string `json:"name"` // Name of the category
	utils.ModelTimestamp        // Embeds the ModelTimestamp struct for timestamp fields
}

// AutoMigrate migrates the Company table if it doesn't exist in the database.
// It uses the provided GORM transaction (`tx`) to perform the migration.
func (v *Company) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&v) {
		err = tx.AutoMigrate(&v)
		if err != nil {
			return err
		}
	}
	return
}

func (v *CompanyCategory) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&v) {
		err = tx.AutoMigrate(&v)
		if err != nil {
			return err
		}
	}
	return
}

// BeforeCreate is a callback function that is executed before creating a new Company record in the database.
// It uses the provided GORM transaction (`tx`) to perform any necessary operations.
// The function calls the BeforeCreate method of the UUIDPrimaryKey field in the Company struct.
// This allows any custom logic or actions defined in the BeforeCreate method of the UUIDPrimaryKey field to be executed before creating the Company record.

func (v *Company) BeforeCreate(tx *gorm.DB) (err error) {
	return v.UUIDPrimaryKey.BeforeCreate(tx)
}

func (v *CompanyCategory) BeforeCreate(tx *gorm.DB) (err error) {
	return v.UUIDPrimaryKey.BeforeCreate(tx)
}
