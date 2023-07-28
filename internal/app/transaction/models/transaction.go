package models

import (
	"backend/internal/app/product/models"
	"backend/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Transaction Product represents a service.product entity.
type Transaction struct {
	utils.UUIDPrimaryKey
	TransactionDate time.Time `json:"transaction_date,omitempty"`
	utils.ModelTimestamp
}

type TransactionDetail struct {
	utils.UUIDPrimaryKey
	TransactionID uuid.UUID      `gorm:"type:uuid"`
	Transactions  Transaction    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TransactionID;references:ID"`
	ProductID     uuid.UUID      `gorm:"type:uuid"`
	Products      models.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductID;references:ID"`
	Qty           int64          `json:"qty"`
	utils.ModelTimestamp
}

// AutoMigrate migrates the Product table if it doesn't exist in the database.
// It uses the provided GORM service.transaction (`tx`) to perform the migrations.
func (v *Transaction) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&v) {
		err = tx.AutoMigrate(&v)
		if err != nil {
			return err
		}
	}
	return
}

func (v *TransactionDetail) AutoMigrate(tx *gorm.DB) (err error) {
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

func (v *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	return v.UUIDPrimaryKey.BeforeCreate(tx)
}

func (v *TransactionDetail) BeforeCreate(tx *gorm.DB) (err error) {
	return v.UUIDPrimaryKey.BeforeCreate(tx)
}
