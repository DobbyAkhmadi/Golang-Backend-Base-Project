package service

import (
	"backend/pkg/utils"
)

type RegionService interface {
	GetPaginationVillage(paginate utils.SetPaginationDto) (utils.GetGlobalResponsePaginationDto, error)
}
