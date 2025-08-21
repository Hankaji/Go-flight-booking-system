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
	ID                  string
	DepartureTime       time.Time
	ArrivalTime         time.Time
	DepartureLocationID string
	ArrivalLocationID   string
	Status              FlightStatus
}

type FlightWithLocation struct {
	ID                string
	DepartureTime     time.Time
	ArrivalTime       time.Time
	DepartureLocation Location `gorm:"embedded"`
	ArrivalLocation   Location `gorm:"embedded"`
	Price             uint
	Airplane          string
	Status            FlightStatus
}
