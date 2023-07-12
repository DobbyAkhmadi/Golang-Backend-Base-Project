package models

import (
	"backend/pkg/utils"
	"gorm.io/gorm"
)

// Product represents a service.product entity.
type Product struct {
	utils.UUIDPrimaryKey        // Embeds the UUIDPrimaryKey struct for a UUID primary key field
	Name                 string `json:"name,omitempty"`        // Name of the service.product
	Description          string `json:"description,omitempty"` // Description of the service.product (optional)
	Stock                int64  `json:"stock"`                 // Stock the service.product
	utils.ModelTimestamp        // Embeds the ModelTimestamp struct for timestamp fields
}

// AutoMigrate migrates the Product table if it doesn't exist in the database.
// It uses the provided GORM service.transaction (`tx`) to perform the migration.
func (v *Product) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&v) {
		err = tx.AutoMigrate(&v)
		if err != nil {
			return err
		}
	}
	return
}

// BeforeCreate is a callback function that is executed before creating a new Product record in the database.
// It uses the provided GORM service.transaction (`tx`) to perform any necessary operations.
// The function calls the BeforeCreate method of the UUIDPrimaryKey field in the Product struct.
// This allows any custom logic or actions defined in the BeforeCreate method of the UUIDPrimaryKey field to be executed before creating the Product record.

func (v *Product) BeforeCreate(tx *gorm.DB) (err error) {
	return v.UUIDPrimaryKey.BeforeCreate(tx)
}
