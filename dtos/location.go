// Package dtos
package dtos

type LocationResponse struct {
	ID           uint    `json:"id"`
	LocationName string  `json:"locationName"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
}
