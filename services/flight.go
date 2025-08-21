// Package services
package services

import (
	"flight-booking-server/dtos"
	"flight-booking-server/repositories"
)

type FlightService struct {
	Repo repositories.FlightRepo
}

// var FlightNotFoundErr = errors.New("Couldnt find a flight")

func (s FlightService) GetAllFlights() ([]dtos.FlightWithLocationResponse, error) {
	flights, err := s.Repo.GetAllWithLocation()
	if err != nil {
		return nil, err
	}

	flightRes := make([]dtos.FlightWithLocationResponse, 0, len(flights))
	for _, flight := range flights {
		flightRes = append(flightRes, dtos.FlightWithLocationResponse{
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
			Status:            flight.Status,
		})
	}

	return flightRes, nil
}

func (s FlightService) GetFlightByID(id string) (*dtos.FlightWithLocationResponse, error) {
	flight, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	flightRes := dtos.FlightWithLocationResponse{
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
		Status: flight.Status,
	}

	return &flightRes, nil
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
