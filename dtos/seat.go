// Package dtos
package dtos

import "flight-booking-server/entities"

type SeatResponse struct {
	ID         uint   `json:"id"`
	PlaneID    uint   `json:"planeId"`
	SeatNumber string `json:"seatNumber"`
}

type SeatAvailabilityResponse struct {
	ID         uint                `json:"id"`
	PlaneID    uint                `json:"planeId"`
	SeatNumber string              `json:"seatNumber"`
	Status     entities.SeatStatus `json:"status"`
}

func SeatE2R(seat *entities.Seat) *SeatResponse {
	if seat == nil {
		return nil
	}

	seatRes := SeatResponse{
		ID:         seat.ID,
		PlaneID:    seat.PlaneID,
		SeatNumber: seat.SeatNumber,
	}

	return &seatRes
}

func SeatAvailabilityE2R(seat *entities.SeatWithStatus) *SeatAvailabilityResponse {
	seatRes := SeatAvailabilityResponse{
		ID:         seat.ID,
		PlaneID:    seat.PlaneID,
		SeatNumber: seat.SeatNumber,
		Status:     seat.Status,
	}

	return &seatRes
}
