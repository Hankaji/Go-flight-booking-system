// Package dtos
package dtos

import "flight-booking-server/entities"

type PlaneResponse struct {
	ID         uint   `json:"id"`
	Model      string `json:"model"`
	TotalSeats uint   `json:"totalSeats"`
}

func PlaneE2R(plane *entities.Plane) *PlaneResponse {
	planeRes := PlaneResponse{
		ID:         plane.ID,
		Model:      plane.Model,
		TotalSeats: plane.TotalSeats,
	}

	return &planeRes
}
