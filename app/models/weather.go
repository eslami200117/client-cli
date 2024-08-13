package model

type Weather struct {
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	WindSpeed   float64 `json:"windSpeed"`
	Temperature float64 `json:"temperature"`
	Rain        float64 `json:"rain"`
}
