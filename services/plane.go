// Package services
package services

import (
	"flight-booking-server/dtos"
	"flight-booking-server/entities"
	"flight-booking-server/repositories"
)

type PlaneService struct {
	Repo repositories.PlaneRepo
}

func (s *PlaneService) GetAllPlanes() ([]dtos.PlaneResponse, error) {
	planes, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	locationsRes := make([]dtos.PlaneResponse, 0, len(planes))
	for _, plane := range planes {
		locationsRes = append(locationsRes, *planeE2R(&plane))
	}

	return locationsRes, nil
}

func (s *PlaneService) GetPlanByID(id string) (*dtos.PlaneResponse, error) {
	plane, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	planeRes := planeE2R(plane)

	return planeRes, nil
}

func planeE2R(plane *entities.Plane) *dtos.PlaneResponse {
	planeRes := dtos.PlaneResponse{
		ID:         plane.ID,
		Model:      plane.Model,
		TotalSeats: plane.TotalSeats,
	}

	return &planeRes
}
