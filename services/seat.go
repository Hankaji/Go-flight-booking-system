// Package services
package services

import (
	"flight-booking-server/dtos"
	"flight-booking-server/repositories"
)

type SeatService struct {
	Repo repositories.SeatRepo
}

func (s *SeatService) GetAllSeats() ([]dtos.SeatResponse, error) {
	seats, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	SeatRes := make([]dtos.SeatResponse, 0, len(seats))
	for _, seat := range seats {
		SeatRes = append(SeatRes, *dtos.SeatE2R(&seat))
	}

	return SeatRes, nil
}

func (s *SeatService) GetSeatByID(id string) (*dtos.SeatResponse, error) {
	seat, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	planeRes := dtos.SeatE2R(seat)

	return planeRes, nil
}
