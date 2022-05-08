package models

import (
	"github.com/google/uuid"
	"github.com/hojabri/backend/pkg/entity"
	"gorm.io/gorm"
)

type Company struct {
	entity.UUIDPrimaryKey
	Name        string           `json:"name" gorm:"uniqueIndex"`
	Description string           `json:"description,omitempty"`
	CategoryID  *uuid.UUID       `json:"category_id,omitempty"`
	Category    *CompanyCategory `json:"category"`
	Website     string           `json:"website,omitempty"`
	entity.ModelTimestamp
}

type CompanyCategory struct {
	entity.UUIDPrimaryKey
	Name string `json:"name" gorm:"uniqueIndex"`
	entity.ModelTimestamp
}

func (v *Company) BeforeCreate(tx *gorm.DB) (err error) {
	return v.UUIDPrimaryKey.BeforeCreate(tx)
}
func (v *CompanyCategory) BeforeCreate(tx *gorm.DB) (err error) {
	return v.UUIDPrimaryKey.BeforeCreate(tx)
}
