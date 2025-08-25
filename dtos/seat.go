// Package dtos
package dtos

import "flight-booking-server/entities"

type SeatResponse struct {
	ID         uint   `json:"id"`
	PlaneID    uint   `json:"planeId"`
	SeatNumber string `json:"seatNumber"`
}

func SeatE2R(seat *entities.Seat) *SeatResponse {
	seatRes := SeatResponse{
		ID:         seat.ID,
		PlaneID:    seat.PlaneID,
		SeatNumber: seat.SeatNumber,
	}

	return &seatRes
}
