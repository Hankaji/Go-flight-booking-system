// Package services
package services

import (
	"errors"
	"flight-booking-server/dtos"
	"flight-booking-server/entities"
	httperrors "flight-booking-server/http-errors"
	"flight-booking-server/repositories"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

type FlightService struct {
	Repo repositories.FlightRepo
}

func (s FlightService) GetAllFlights(pagination *repositories.Pagination, filter *repositories.FlightFilter) ([]dtos.FlightResponse, error) {
	flights, err := s.Repo.GetAllWithLocationWithParams(pagination, filter)
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

func (s FlightService) GetFlightByID(id uint) (*dtos.FlightResponse, error) {
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

func (s FlightService) GetAllTickets(flightID string) ([]dtos.TicketResponse, error) {
	tickets, err := s.Repo.GetTickets(flightID, nil)
	if err != nil {
		return nil, err
	}

	ticketRes := make([]dtos.TicketResponse, 0, len(tickets))
	for _, ticket := range tickets {
		ticketRes = append(ticketRes, *dtos.TicketE2R(&ticket))
	}

	return ticketRes, nil
}

func (s FlightService) CreateFlight(req dtos.CreateFlightRequest) error {
	// Validate time
	var flights []entities.Flight
	{
		_flights, err := s.Repo.GetAllWithParams(nil, &repositories.FlightFilter{
			PlaneID: &req.AirplaneID,
		})
		if err != nil {
			return err
		}
		flights = _flights
	}

	println("Filtered flights: ", flights)

	// Validate location

	_, err := s.Repo.CreateFlight(req)
	if err != nil {
		return err
	}

	return nil
}

func (s FlightService) CreateTicket(flightID uint, req dtos.CreateTicketRequest) error {
	_, err := s.Repo.CreateTicket(flightID, req)
	if err != nil {
		return err
	}

	return nil
}

func (s FlightService) UpdateTicket(flightID, ticketID uint, req dtos.UpdateTicketRequest) error {
	// Validate flight exists
	if _, err := s.Repo.GetByID(flightID); err != nil {
		return err
	}

	// Validate existingTicket in flight
	if _, err := s.Repo.GetTicketByID(flightID, ticketID); err != nil {
		return err
	}

	updates := map[string]any{}
	if req.SeatID != nil {
		updates["seat_id"] = *req.SeatID
	}
	if req.Username != nil {
		updates["username"] = *req.Username
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if len(updates) > 0 {
		err := s.Repo.UpdateTicket(flightID, ticketID, updates)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s FlightService) UpdateFlight(flightID uint, req dtos.UpdateFlightRequest) error {
	// Validate flight exists
	if _, err := s.Repo.GetByID(flightID); err != nil {
		return err
	}

	// check if there is any tickets booked on this
	approvedTicketStatus := entities.TicketApproved
	if flightTickets, err := s.Repo.GetTickets(strconv.FormatUint(uint64(flightID), 10), &repositories.TicketFilter{
		Status: &approvedTicketStatus,
	}); err != nil {
		return err
	} else {
		if len(flightTickets) != 0 && req.PlaneID != nil {
			return httperrors.NewHTTPErr(http.StatusConflict,
				fmt.Errorf("couldn't update plane on flight ID %d as there are still tickets booked on this", flightID),
				nil,
			)
		}
	}

	updates := map[string]any{}
	if req.DepartureTime != nil {
		updates["departure_time"] = *req.DepartureTime
	}
	if req.ArrivalTime != nil {
		updates["arrival_time"] = *req.ArrivalTime
	}
	if req.DepartureLocationID != nil {
		updates["departure_location_id"] = *req.DepartureLocationID
	}
	if req.ArrivalLocationID != nil {
		updates["arrival_location_id"] = *req.ArrivalLocationID
	}
	if req.PlaneID != nil {
		updates["plane_id"] = *req.PlaneID
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	err := s.Repo.UpdateFlight(flightID, updates)
	if err != nil {
		return err
	}

	return nil
}

func (s FlightService) DeleteFlight(flightID string) error {
	// Validate if flight can be deleted
	approvedTicketStatus := entities.TicketApproved
	if flightTickets, err := s.Repo.GetTickets(flightID, &repositories.TicketFilter{
		Status: &approvedTicketStatus,
	}); err != nil {
		return err
	} else {
		if len(flightTickets) != 0 {
			return httperrors.NewHTTPErr(http.StatusConflict,
				errors.New("flight couldn't be deleted as there are still tickets booked on this"),
				nil)
		}
	}

	err := s.Repo.DeleteFlight(flightID)
	if err != nil {
		return err
	}

	return nil
}
