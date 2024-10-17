package service

import (
	"encoding/json"
	"fmt"
	"github.com/ARTM2000/rahgozar/internal/core/dto"
	"github.com/ARTM2000/rahgozar/internal/core/port"
	"log/slog"
	"os"
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

func (mls *mapLayersService) GetActiveMapLayerByName(layerName string) (dto.MapLayerFullInfo, error) {
	points := mls.readDataFromJSONFiles(layerName)
	return dto.MapLayerFullInfo{
		MapLayerCompactInfo: dto.MapLayerCompactInfo{ID: 1, Name: "subway", Title: "مترو", Image: ""},
		Points:              *points,
		Lines:               []dto.GeoJSON[dto.GeoJSONLineStringFeature]{},
	}, nil
}

func (mls *mapLayersService) readDataFromJSONFiles(layerName string) *[]dto.GeoJSON[dto.GeoJSONPointFeature] {
	var points []dto.GeoJSON[dto.GeoJSONPointFeature]
	dataDirectory := fmt.Sprintf("./data/%s", layerName)
	dir, err := os.ReadDir(dataDirectory)
	if err != nil {
		slog.Error("error reading data", slog.Any("err", err))
		return nil
	}
	for _, d := range dir {
		info, _ := d.Info()
		fileName := fmt.Sprintf("%s/%s", dataDirectory, info.Name())
		file, _ := os.ReadFile(fileName)
		var d dto.GeoJSON[dto.GeoJSONPointFeature]
		_ = json.Unmarshal(file, &d)
		points = append(points, d)
		break
	}

	return &points
}
