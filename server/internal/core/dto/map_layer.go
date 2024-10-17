package dto

type MapLayerCompactInfo struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type GeoJSONGeometry[T any] struct {
	Type        string `json:"type"`
	Coordinates T      `json:"coordinates"`
}

type geoJSON[T any] struct {
	Type       string                 `json:"type"`
	Geometry   GeoJSONGeometry[T]     `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type GeoJSONPointFeature struct {
	geoJSON[[]float64]
}

type GeoJSONLineStringFeature struct {
	geoJSON[[][]float64]
}

type MapLayerFullInfo struct {
	MapLayerCompactInfo
	Points []GeoJSONPointFeature      `json:"points"`
	Lines  []GeoJSONLineStringFeature `json:"lines"`
}
