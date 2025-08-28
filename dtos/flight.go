// Package dtos for apis
package dtos

import (
	"flight-booking-server/entities"
	"time"
)

type LocationSummaryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FlightResponse struct {
	ID                uint                    `json:"id"`
	DepartureTime     time.Time               `json:"departureTime"`
	ArrivalTime       time.Time               `json:"arrivalTime"`
	DepartureLocation LocationSummaryResponse `json:"departureLocation"`
	ArrivalLocation   LocationSummaryResponse `json:"arrivalLocation"`
	EstimatedDuration time.Duration           `json:"estimatedDuration"`
	Price             float32                 `json:"price" binding:"required,min=1"`
	Airplane          PlaneResponse           `json:"airplane" binding:"required,min=1,max=50"`
	Status            entities.FlightStatus   `json:"status"`
}

type CreateFlightRequest struct {
	DepartureTime       time.Time `json:"departureTime" binding:"required"`
	ArrivalTime         time.Time `json:"arrivalTime" binding:"required"`
	DepartureLocationID string    `json:"departureLocationID" binding:"required"`
	ArrivalLocationID   string    `json:"arrivalLocationID" binding:"required"`
	Price               float32   `json:"price" binding:"required,min=1"`
	Airplane            string    `json:"airplane" binding:"required,min=1,max=50"`
}

type CreateTicketRequest struct {
	SeatID   uint   `json:"seatId" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UpdateTicketRequest struct {
	SeatID   *uint                  `json:"seatId,omitempty"`
	Username *string                `json:"username,omitempty"`
	Status   *entities.TicketStatus `json:"status,omitempty" binding:"omitempty,oneof=approved cancelled"`
}

type UpdateFlightRequest struct {
	DepartureTime       *time.Time             `json:"departureTime,omitempty"`
	ArrivalTime         *time.Time             `json:"arrivalTime,omitempty"`
	DepartureLocationID *uint                  `json:"departureLocationId,omitempty"`
	ArrivalLocationID   *uint                  `json:"arrivalLocationId,omitempty"`
	PlaneID             *uint                  `json:"planeId,omitempty"`
	Status              *entities.FlightStatus `json:"status" binding:"omitempty,oneof=scheduled delayed departed arrived cancelled"`
}
