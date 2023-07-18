package repository

import (
	models2 "backend/internal/app/region/models"
	"backend/pkg/utils"
	"backend/platform/seeds"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// RegionRepository is an interface that defines the contract for accessing and manipulating Region data.
type RegionRepository interface {
	GetPaginationVillage(paginate utils.SetPaginationDto) ([]*models2.Village, int64, error)
}

// dbRegionRepository implements the RegionRepository interface.
type dbRegionRepository struct {
	connection *gorm.DB
}

// NewRegionRepository creates a new instance of the RegionRepository.
// It performs necessary database migrations and generates fake data if in the development environment.
func NewRegionRepository() RegionRepository {
	// Check if the database connection is already established
	if utils.DB == nil {
		// Connect to the database
		database, _ := utils.Connect()
		if database != nil {
			log.Error(database)
		}
	}

	// Perform auto-migration for Region table
	mProvince := models2.Province{}
	pErr := mProvince.AutoMigrate(utils.DB)
	if pErr != nil {
		panic(pErr)
	}

	// Perform auto-migration for Region table
	mRegency := models2.Regency{}
	rErr := mRegency.AutoMigrate(utils.DB)
	if rErr != nil {
		panic(rErr)
	}

	// Perform auto-migration for Region table
	mDistrict := models2.District{}
	dErr := mDistrict.AutoMigrate(utils.DB)
	if dErr != nil {
		panic(dErr)
	}

	// Perform auto-migration for Region table
	mVillage := models2.Village{}
	vErr := mVillage.AutoMigrate(utils.DB)
	if vErr != nil {
		panic(vErr)
	}

	log.Info("Migration completed successfully!")

	seeds.MainSeed("region")
	// Return the dbRegionRepository instance
	return &dbRegionRepository{
		connection: utils.DB,
	}
}

func (db dbRegionRepository) GetPaginationVillage(paginate utils.SetPaginationDto) ([]*models2.Village, int64, error) {
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
