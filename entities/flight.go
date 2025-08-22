// Package entities represent all the entities in the application
package entities

import "time"

type FlightStatus string

const (
	FlightScheduled FlightStatus = "scheduled"
	FlightDelayed   string       = "delayed"
	FlightDeparted               = "departed"
	FlightArrived                = "arrived"
	FlightCancelled              = "cancelled"
)

type Flight struct {
	ID                  int
	DepartureTime       time.Time
	ArrivalTime         time.Time
	DepartureLocationID string
	DepartureLocation   Location `gorm:"foreignKey:DepartureLocationID;references:ID"`

	ArrivalLocationID string
	ArrivalLocation   Location `gorm:"foreignKey:ArrivalLocationID;references:ID"`

	Price   float32
	PlaneID string
	Plane   Plane `gorm:"foreignKey:PlaneID;references:ID"`

	Status FlightStatus
}

type FlightDetailed struct {
	ID                int
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
