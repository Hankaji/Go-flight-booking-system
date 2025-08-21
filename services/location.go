// Package services
package services

import (
	"flight-booking-server/dtos"
	"flight-booking-server/repositories"
)

type LocationService struct {
	Repo repositories.LocationRepo
}

func (s *LocationService) GetAllLocations() ([]dtos.LocationResponse, error) {
	locations, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	locationsRes := make([]dtos.LocationResponse, 0, len(locations))
	for _, loc := range locations {
		locationsRes = append(locationsRes, dtos.LocationResponse{
			ID:           loc.ID,
			LocationName: loc.LocationName,
			Latitude:     loc.Latitude,
			Longitude:    loc.Longitude,
		})
	}

	return locationsRes, nil
}

func (s *LocationService) GetLocationByID(id string) (*dtos.LocationResponse, error) {
	loc, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	locationRes := dtos.LocationResponse{
		ID:           loc.ID,
		LocationName: loc.LocationName,
		Latitude:     loc.Latitude,
		Longitude:    loc.Longitude,
	}

	return &locationRes, nil
}
