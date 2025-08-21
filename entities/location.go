// Package entities
package entities

type Location struct {
	ID           string `gorm:"primaryKey"`
	LocationName string
	Latitude     float32
	Longitude    float32
}
