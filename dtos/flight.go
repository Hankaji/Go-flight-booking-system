// Package dtos for apis
package dtos

import (
	flight "flight-booking-server/entities"
	"time"
)

type LocationSummaryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FlightResponse struct {
	ID                string                  `json:"id"`
	DepartureTime     time.Time               `json:"departureTime"`
	ArrivalTime       time.Time               `json:"arrivalTime"`
	DepartureLocation LocationSummaryResponse `json:"departureLocation"`
	ArrivalLocation   LocationSummaryResponse `json:"arrivalLocation"`
	Status            flight.FlightStatus     `json:"status"`
}

type FlightWithLocationResponse struct {
	ID                string                  `json:"id"`
	DepartureTime     time.Time               `json:"departureTime"`
	ArrivalTime       time.Time               `json:"arrivalTime"`
	DepartureLocation LocationSummaryResponse `json:"departureLocation"`
	ArrivalLocation   LocationSummaryResponse `json:"arrivalLocation"`
	EstimatedDuration time.Duration           `json:"estimatedDuration"`
	Price             uint                    `json:"price" binding:"required,min=1"`
	Airplane          string                  `json:"airplane" binding:"required,min=1,max=50"`
	Status            flight.FlightStatus     `json:"status"`
}

type CreateFlightRequest struct {
	DepartureTime       time.Time `json:"departureTime" binding:"required"`
	ArrivalTime         time.Time `json:"arrivalTime" binding:"required"`
	DepartureLocationID string    `json:"departureLocationID" binding:"required"`
	ArrivalLocationID   string    `json:"arrivalLocationID" binding:"required"`
	Price               int       `json:"price" binding:"required,min=1"`
	Airplane            string    `json:"airplane" binding:"required,min=1,max=50"`
}

type UpdateFlightRequest struct {
	DepartureTime     *time.Time           `json:"departureTime,omitempty"`
	ArrivalTime       *time.Time           `json:"arrivalTime,omitempty"`
	DepartureLocation *string              `json:"departureLocation,omitempty" binding:"omitempty,min=3,max=100"`
	ArrivalLocation   *string              `json:"arrivalLocation,omitempty" binding:"omitempty,min=3,max=100"`
	Status            *flight.FlightStatus `json:"status" binding:"required,oneof=scheduled delayed departed arrived cancelled"`
}
