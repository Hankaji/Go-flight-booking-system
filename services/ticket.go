// Package services
package services

import (
	"flight-booking-server/dtos"
	"flight-booking-server/repositories"
)

type TicketService struct {
	Repo repositories.TicketRepo
}

func (s *TicketService) GetAllTickets() ([]dtos.TicketResponse, error) {
	tickets, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	ticketsRes := make([]dtos.TicketResponse, 0, len(tickets))
	for _, ticket := range tickets {
		ticketsRes = append(ticketsRes, *dtos.TicketE2R(&ticket))
	}

	return ticketsRes, nil
}

func (s *TicketService) GetTicketByID(id string) (*dtos.TicketResponse, error) {
	ticket, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	locationRes := dtos.TicketE2R(ticket)

	return locationRes, nil
}

func (s *TicketService) CreateTicket(req dtos.CreateTicketRequest) error {
	_, err := s.Repo.CreateTicket(req)
	if err != nil {
		return err
	}

	return nil
}
