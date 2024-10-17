package port

import "github.com/ARTM2000/rahgozar/internal/core/dto"

type IMapLayersService interface {
	GetActiveMapLayersList() ([]dto.MapLayerCompactInfo, error)

	GetActiveMapLayerByName(layerName string) (dto.MapLayerFullInfo, error)
}
