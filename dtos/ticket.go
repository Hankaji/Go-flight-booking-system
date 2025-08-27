// Package dtos
package dtos

import "flight-booking-server/entities"

type TicketResponse struct {
	ID       uint                  `json:"id"`
	SeatID   uint                  `json:"seatId"`
	Username string                `json:"username"`
	Status   entities.TicketStatus `json:"status"`

	Seat *SeatResponse `json:"seat,omitempty"`
}

func TicketE2R(ticket *entities.Ticket) *TicketResponse {
	if ticket == nil {
		return nil
	}

	ticketRes := TicketResponse{
		ID:       ticket.ID,
		SeatID:   ticket.SeatID,
		Username: ticket.Username,
		Status:   ticket.Status,

		Seat: SeatE2R(ticket.Seat),
	}

	return &ticketRes
}
