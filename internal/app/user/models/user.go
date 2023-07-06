package models

import (
	"backend/pkg/entity"
	"gorm.io/gorm"
)

type User struct {
	entity.UUIDPrimaryKey
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
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
