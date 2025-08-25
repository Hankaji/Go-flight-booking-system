// Package services
package services

import (
	"flight-booking-server/dtos"
	"flight-booking-server/repositories"
	"sort"
)

type FlightService struct {
	Repo repositories.FlightRepo
}

// var FlightNotFoundErr = errors.New("Couldnt find a flight")

func (s FlightService) GetAllFlights() ([]dtos.FlightResponse, error) {
	flights, err := s.Repo.GetAllWithLocation()
	if err != nil {
		return nil, err
	}

	flightRes := make([]dtos.FlightResponse, 0, len(flights))
	for _, flight := range flights {
		flightRes = append(flightRes, dtos.FlightResponse{
			ID:            flight.ID,
			DepartureTime: flight.DepartureTime,
			ArrivalTime:   flight.ArrivalTime,
			DepartureLocation: dtos.LocationSummaryResponse{
				ID:   flight.DepartureLocation.ID,
				Name: flight.DepartureLocation.LocationName,
			},
			ArrivalLocation: dtos.LocationSummaryResponse{
				ID:   flight.ArrivalLocation.ID,
				Name: flight.ArrivalLocation.LocationName,
			},
			EstimatedDuration: flight.ArrivalTime.Sub(flight.DepartureTime),
			Airplane:          *dtos.PlaneE2R(&flight.Plane),
			Price:             flight.Price,
			Status:            flight.Status,
		})
	}

	return flightRes, nil
}

func (s FlightService) GetFlightByID(id string) (*dtos.FlightResponse, error) {
	flight, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	flightRes := dtos.FlightResponse{
		ID:            flight.ID,
		DepartureTime: flight.DepartureTime,
		ArrivalTime:   flight.ArrivalTime,
		DepartureLocation: dtos.LocationSummaryResponse{
			ID:   flight.DepartureLocation.ID,
			Name: flight.DepartureLocation.LocationName,
		},
		ArrivalLocation: dtos.LocationSummaryResponse{
			ID:   flight.ArrivalLocation.ID,
			Name: flight.ArrivalLocation.LocationName,
		},
		Airplane: *dtos.PlaneE2R(&flight.Plane),

		Status: flight.Status,
	}

	return &flightRes, nil
}

func (s FlightService) GetAllSeatsAvailability(flightID string) ([]dtos.SeatAvailabilityResponse, error) {
	seats, err := s.Repo.GetSeatAvailability(flightID)
	if err != nil {
		return nil, err
	}

	// Ascending
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].ID < seats[j].ID
	})

	SeatRes := make([]dtos.SeatAvailabilityResponse, 0, len(seats))
	for _, seat := range seats {
		SeatRes = append(SeatRes, *dtos.SeatAvailabilityE2R(&seat))
	}

	return SeatRes, nil
}

func (s FlightService) CreateFlight(req dtos.CreateFlightRequest) error {
	// Validate time
	// Validate location

	_, err := s.Repo.CreateFlight(req)
	if err != nil {
		return err
	}

	return nil
}
