package repository

import (
	models2 "backend/internal/app/region/models"
	"backend/pkg/utils"
	"backend/platform/database"
	"backend/platform/seeders"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// RegionRepository is an interface that defines the contract for accessing and manipulating Region data.
type RegionRepository interface {
	SaveV(model *models2.Village) (*models2.Village, error)
	SaveP(model *models2.Province) (*models2.Province, error)
	SaveD(model *models2.District) (*models2.District, error)
	SaveR(model *models2.Regency) (*models2.Regency, error)
	GetPaginationV(paginate utils.SetPaginationDto) ([]*models2.Village, int64, error)
	GetPaginationP(paginate utils.SetPaginationDto) ([]*models2.Province, int64, error)
	GetPaginationD(paginate utils.SetPaginationDto) ([]*models2.District, int64, error)
	GetPaginationR(paginate utils.SetPaginationDto) ([]*models2.Regency, int64, error)
}

// dbRegionRepository implements the RegionRepository interface.
type dbRegionRepository struct {
	connection *gorm.DB
}

// NewRegionRepository creates a new instance of the RegionRepository.
// It performs necessary database migrations and generates fake data if in the development environment.
func NewRegionRepository() RegionRepository {
	// Check if the database connection is already established
	if database.DB == nil {
		// Connect to the database
		database, _ := database.Connect()
		if database != nil {
			log.Error(database)
		}
	}

	// Perform auto-migrations for Region table
	mProvince := models2.Province{}
	pErr := mProvince.AutoMigrate(database.DB)
	if pErr != nil {
		panic(pErr)
	}

	// Perform auto-migrations for Region table
	mRegency := models2.Regency{}
	rErr := mRegency.AutoMigrate(database.DB)
	if rErr != nil {
		panic(rErr)
	}

	// Perform auto-migrations for Region table
	mDistrict := models2.District{}
	dErr := mDistrict.AutoMigrate(database.DB)
	if dErr != nil {
		panic(dErr)
	}

	// Perform auto-migrations for Region table
	mVillage := models2.Village{}
	vErr := mVillage.AutoMigrate(database.DB)
	if vErr != nil {
		panic(vErr)
	}

	log.Info("Migration completed successfully!")

	seeders.MainSeed("region")
	// Return the dbRegionRepository instance
	return &dbRegionRepository{
		connection: database.DB,
	}
}

func (db dbRegionRepository) GetPaginationV(paginate utils.SetPaginationDto) ([]*models2.Village, int64, error) {
	offset := (paginate.PageIndex - 1) * paginate.PageSize
	var myModel []*models2.Village
	var total int64

	query := db.connection.Model(&models2.Village{})

	// Apply sorting
	if paginate.SortBy != "" {
		query = query.Order(paginate.SortBy)
	}

	// Apply global search
	if paginate.GlobalSearch != "" {
		search := "%" + paginate.GlobalSearch + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", search, search)
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
func (db dbRegionRepository) GetPaginationP(paginate utils.SetPaginationDto) ([]*models2.Province, int64, error) {
	offset := (paginate.PageIndex - 1) * paginate.PageSize
	var myModel []*models2.Province
	var total int64

	query := db.connection.Model(&models2.Province{})

	// Apply sorting
	if paginate.SortBy != "" {
		query = query.Order(paginate.SortBy)
	}

	// Apply global search
	if paginate.GlobalSearch != "" {
		search := "%" + paginate.GlobalSearch + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", search, search)
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
func (db dbRegionRepository) GetPaginationD(paginate utils.SetPaginationDto) ([]*models2.District, int64, error) {
	offset := (paginate.PageIndex - 1) * paginate.PageSize
	var myModel []*models2.District
	var total int64

	query := db.connection.Model(&models2.District{})

	// Apply sorting
	if paginate.SortBy != "" {
		query = query.Order(paginate.SortBy)
	}

	// Apply global search
	if paginate.GlobalSearch != "" {
		search := "%" + paginate.GlobalSearch + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", search, search)
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
func (db dbRegionRepository) GetPaginationR(paginate utils.SetPaginationDto) ([]*models2.Regency, int64, error) {
	offset := (paginate.PageIndex - 1) * paginate.PageSize
	var myModel []*models2.Regency
	var total int64

	query := db.connection.Model(&models2.Regency{})

	// Apply sorting
	if paginate.SortBy != "" {
		query = query.Order(paginate.SortBy)
	}

	// Apply global search
	if paginate.GlobalSearch != "" {
		search := "%" + paginate.GlobalSearch + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", search, search)
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

func (db dbRegionRepository) SaveV(model *models2.Village) (*models2.Village, error) {
	if err := db.connection.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
func (db dbRegionRepository) SaveP(model *models2.Province) (*models2.Province, error) {
	if err := db.connection.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
func (db dbRegionRepository) SaveD(model *models2.District) (*models2.District, error) {
	if err := db.connection.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
func (db dbRegionRepository) SaveR(model *models2.Regency) (*models2.Regency, error) {
	if err := db.connection.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
