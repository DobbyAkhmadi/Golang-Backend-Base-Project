package repository

import (
	"backend/internal/app/user/models"
	"backend/pkg/entity"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Save(user models.User) (uuid.UUID, error)
	Update(models.User) error
	Delete(models.User) error
	FindAll() []*models.User
	FindByID(userID uuid.UUID) (*models.User, error)
	DeleteByID(userID uuid.UUID) error
	FindByName(name string) (*models.User, error)
	FindByField(fieldName, fieldValue string) (*models.User, error)
	UpdateSingleField(user models.User, fieldName, fieldValue string) error
}
type userDatabase struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	if entity.DB == nil {
		connect, _ := entity.Connect()
		if connect != nil {
			log.Error(connect)
		}
		model := models.User{}
		err := model.AutoMigrate(entity.DB)
		if err != nil {
			panic(err)
		}
	}
	return &userDatabase{
		connection: entity.DB,
	}
}

func (db userDatabase) DeleteByID(userID uuid.UUID) error {
	user := models.User{}
	user.ID = userID
	result := db.connection.Delete(&user)
	return result.Error
}

func (db userDatabase) Save(user models.User) (uuid.UUID, error) {
	result := db.connection.Create(&user)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return user.ID, nil
}

func (db userDatabase) Update(user models.User) error {
	result := db.connection.Save(&user)
	return result.Error
}

func (db userDatabase) Delete(user models.User) error {
	result := db.connection.Delete(&user)
	return result.Error
}

func (db userDatabase) FindAll() []*models.User {
	var users []*models.User
	db.connection.Preload(clause.Associations).Find(&users)
	return users
}

func (db userDatabase) FindByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	result := db.connection.Preload(clause.Associations).Find(&user, "id = ?", userID)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

func (db userDatabase) FindByName(name string) (*models.User, error) {
	var user models.User
	result := db.connection.Preload(clause.Associations).Find(&user, "name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

func (db userDatabase) FindByField(fieldName, fieldValue string) (*models.User, error) {
	var user models.User
	result := db.connection.Preload(clause.Associations).Find(&user, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

func (db userDatabase) UpdateSingleField(user models.User, fieldName, fieldValue string) error {
	result := db.connection.Model(&user).Update(fieldName, fieldValue)
	return result.Error
}
