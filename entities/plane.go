// Package entities
package entities

type Plane struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Model      string
	TotalSeats uint

	Seats []Seat `gorm:"foreignKey:PlaneID"`
}
