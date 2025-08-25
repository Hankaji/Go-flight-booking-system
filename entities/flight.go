// Package entities represent all the entities in the application
package entities

import (
	"time"
)

type FlightStatus string

const (
	FlightScheduled FlightStatus = "scheduled"
	FlightDelayed   string       = "delayed"
	FlightDeparted               = "departed"
	FlightArrived                = "arrived"
	FlightCancelled              = "cancelled"
)

type Flight struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	DepartureTime time.Time
	ArrivalTime   time.Time

	DepartureLocationID string
	ArrivalLocationID   string

	Price   float32
	PlaneID string
	Status  FlightStatus

	DepartureLocation Location `gorm:"foreignKey:DepartureLocationID;references:ID"`
	ArrivalLocation   Location `gorm:"foreignKey:ArrivalLocationID;references:ID"`
	Plane             Plane    `gorm:"foreignKey:PlaneID;references:ID"`
}

type FlightDetailed struct {
	ID                uint
	DepartureTime     time.Time
	ArrivalTime       time.Time
	DepartureLocation Location `gorm:"embedded"`
	ArrivalLocation   Location `gorm:"embedded"`
	Price             uint
	Plane             Plane `gorm:"embedded"`
	Status            FlightStatus
}

func (FlightDetailed) TableName() string {
	return "flights"
}
