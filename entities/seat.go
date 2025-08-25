// Package entities
package entities

type Seat struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	PlaneID    uint   `gorm:"not null"`
	SeatNumber string `gorm:"type:varchar(50);not null"`

	Plane Plane `gorm:"foreignKey:PlaneID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
