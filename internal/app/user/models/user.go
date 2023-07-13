package models

import (
	"backend/pkg/utils"
	"gorm.io/gorm"
)

// User represents a user entity.
type User struct {
	utils.UUIDPrimaryKey        // Embeds the UUIDPrimaryKey struct for a UUID primary key field
	FirstName            string `json:"first_name,omitempty"`             // FirstName of the user (optional)
	LastName             string `json:"last_name,omitempty"`              // LastName of the user (optional)
	PhoneNumber          string `json:"phone_number,omitempty"`           // PhoneNumber of the user (optional)
	Address              string `json:"address,omitempty"`                // Address of the user (optional)
	Username             string `json:"username,omitempty"`               // Username of the user (optional)
	Password             string `json:"password" validate:"required"`     // Password of the user (required)
	Email                string `json:"email,omitempty" validate:"email"` // Email of the user (optional, must be a valid email address)
	utils.ModelTimestamp        // Embeds the ModelTimestamp struct for timestamp fields
}

// AutoMigrate migrates the User table if it doesn't exist in the database.
// It uses the provided GORM service.transaction (`tx`) to perform the migration.
func (v *User) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&v) {
		err = tx.AutoMigrate(&v)
		if err != nil {
			return err
		}
	}
	return
}

// BeforeCreate is a callback function that is executed before creating a new User record in the database.
// It uses the provided GORM service.transaction (`tx`) to perform any necessary operations.
// The function calls the BeforeCreate method of the UUIDPrimaryKey field in the User struct.
// This allows any custom logic or actions defined in the BeforeCreate method of the UUIDPrimaryKey field to be executed before creating the User record.
func (v *User) BeforeCreate(tx *gorm.DB) (err error) {
	return v.UUIDPrimaryKey.BeforeCreate(tx)
}
