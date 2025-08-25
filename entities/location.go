// Package entities
package entities

type Location struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	LocationName string
	Latitude     float32
	Longitude    float32
}
