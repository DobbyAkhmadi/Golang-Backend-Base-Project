package service

import (
	"backend/internal/app/region/repository"
	"backend/pkg/utils"
)

type RegionServiceImpl struct {
	regionRepository repository.RegionRepository
}

func NewRegionService(regionRepository *repository.RegionRepository) *RegionServiceImpl {
	return &RegionServiceImpl{
		regionRepository: *regionRepository,
	}
}

func (s *RegionServiceImpl) GetPaginationVillage(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error) {
	// Retrieve paginated product from the repository
	products, total, err := s.regionRepository.GetPaginationVillage(paginate)
	if err != nil {
		return utils.GetGlobalResponsePaginationDto{}, err
	}

	generate := utils.GetGlobalResponsePaginationDto{
		Header: utils.HeaderDto{
			Milliseconds: utils.GetCurrentLatency(),
			Message:      "Request Successfully",
		},
		Code:      200,
		Status:    "OK",
		Data:      products,
		PageIndex: paginate.PageIndex,
		PageSize:  paginate.PageSize,
		TotalRows: total,
	}

	return generate, nil
}
