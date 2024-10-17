package service

import (
	"github.com/ARTM2000/rahgozar/internal/core/dto"
	"github.com/ARTM2000/rahgozar/internal/core/port"
)

type mapLayersService struct{}

func NewMapLayersService() port.IMapLayersService {
	return &mapLayersService{}
}

func (mls *mapLayersService) GetActiveMapLayersList() ([]dto.MapLayerCompactInfo, error) {
	return []dto.MapLayerCompactInfo{
		{ID: 1, Name: "subway", Title: "مترو", Image: ""},
		{ID: 2, Name: "brt", Title: "اتوبوس BRT", Image: ""},
		{ID: 3, Name: "bus", Title: "اتوبوس خطی", Image: ""},
		{ID: 4, Name: "taxi", Title: "تاکسی", Image: ""},
	}, nil
}

func (mls *mapLayersService) GetActiveMapLayerByName(name string) (dto.MapLayerFullInfo, error) {
	return dto.MapLayerFullInfo{
		MapLayerCompactInfo: dto.MapLayerCompactInfo{ID: 1, Name: "subway", Title: "مترو", Image: ""},
		Points:              nil,
		Lines:               nil,
	}, nil
}
