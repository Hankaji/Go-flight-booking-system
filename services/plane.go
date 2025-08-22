// Package services
package services

import (
	"flight-booking-server/dtos"
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
		locationsRes = append(locationsRes, *dtos.PlaneE2R(&plane))
	}

	return locationsRes, nil
}

func (s *PlaneService) GetPlanByID(id string) (*dtos.PlaneResponse, error) {
	plane, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	planeRes := dtos.PlaneE2R(plane)

	return planeRes, nil
}
