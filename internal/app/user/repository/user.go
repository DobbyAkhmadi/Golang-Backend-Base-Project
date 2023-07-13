package repository

import (
	models2 "backend/internal/app/user/models"
	"backend/pkg/utils"
	"gorm.io/gorm"
)

// UserRepository is an interface that defines the contract for accessing and manipulating User data.
type UserRepository interface {
	// Save basic crud
	Save(model *models2.User) (*models2.User, error)
	Update(model *models2.User) (*models2.User, error)
	GetByID(id string) (*models2.User, error)
	Delete(id string) error
	GetPagination(paginate utils.SetPaginationDto) ([]*models2.User, int64, error)
}

// dbUserRepository implements the UserRepository interface.
type dbUserRepository struct {
	connection *gorm.DB
}

func (db *dbUserRepository) Save(model *models2.User) (*models2.User, error) {
	if err := db.connection.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (db *dbUserRepository) Update(model *models2.User) (*models2.User, error) {
	if err := db.connection.Save(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (db *dbUserRepository) GetByID(id string) (*models2.User, error) {
	user := &models2.User{}
	if err := db.connection.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (db *dbUserRepository) Delete(id string) error {
	user := &models2.User{}
	if err := db.connection.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (db *dbUserRepository) GetPagination(paginate utils.SetPaginationDto) ([]*models2.User, int64, error) {
	offset := (paginate.PageIndex - 1) * paginate.PageSize
	var myModel []*models2.User
	var total int64

	query := db.connection.Model(&models2.User{})

	// Apply sorting
	if paginate.SortBy != "" {
		query = query.Order(paginate.SortBy)
	}

	// Apply global search
	if paginate.GlobalSearch != "" {
		search := "%" + paginate.GlobalSearch + "%"
		query = query.Where("first_name LIKE ? OR last_name LIKE ?OR phone_number LIKE ?", search, search, search)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if err := query.Offset(offset).Limit(paginate.PageSize).Find(&myModel).Error; err != nil {
		return nil, 0, err
	}

	return myModel, total, nil
}
