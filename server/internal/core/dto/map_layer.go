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

type GeoJSONFeature[T any] struct {
	Type       string                 `json:"type"`
	Geometry   GeoJSONGeometry[T]     `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type GeoJSONPointFeature struct {
	GeoJSONFeature[[]string]
}

type GeoJSONLineStringFeature struct {
	GeoJSONFeature[[][]string]
}

type GeoJSON[T any] struct {
	Type     string `json:"type"`
	Features []T    `json:"features"`
}

type MapLayerFullInfo struct {
	MapLayerCompactInfo
	Points []GeoJSON[GeoJSONPointFeature]      `json:"points"`
	Lines  []GeoJSON[GeoJSONLineStringFeature] `json:"lines"`
}
