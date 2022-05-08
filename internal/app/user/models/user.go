package models

import (
	"backend/internal/app/company/models"
	"backend/pkg/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	entity.UUIDPrimaryKey
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Email     string          `json:"email,omitempty" gorm:"uniqueIndex"`
	Phone     string          `json:"phone,omitempty"`
	CompanyID *uuid.UUID      `json:"company_id,omitempty"`
	Company   *models.Company `json:"company,omitempty"`
	entity.ModelTimestamp
}

func (u *User) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&u) {
		err = tx.AutoMigrate(&u)
		if err != nil {
			return err
		}
	}
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return u.UUIDPrimaryKey.BeforeCreate(tx)
}
