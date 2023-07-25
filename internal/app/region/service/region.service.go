package service

import (
	"backend/pkg/utils"
)

type RegionService interface {
	GetPaginationVillage(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
	GetPaginationProvince(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
	GetPaginationDistrict(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
	GetPaginationRegency(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
}
