// Package repositories
package repositories

import (
	"errors"
	"flight-booking-server/entities"
	httperrors "flight-booking-server/http-errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type ITicketRepo interface {
	IRepoRead[entities.Ticket]
	IRepoCreate[entities.Ticket]
}

type TicketRepo struct {
	DB *gorm.DB
}

type TicketFilter struct {
	SeatID   *uint
	FlightID *uint
	Username *string
	Status   *entities.TicketStatus
}

func (repo *TicketRepo) GetAll() ([]entities.Ticket, error) {
	db := repo.DB

	var tickets []entities.Ticket
	db.Find(&tickets)

	return tickets, nil
}

func (repo *TicketRepo) GetByID(id string) (*entities.Ticket, error) {
	db := repo.DB

	var ticket entities.Ticket
	err := db.First(&ticket, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("No ticket found with that ID")
		return nil, httperrors.NewHTTPErr(http.StatusNotFound, fmt.Errorf("ticket with ID %s does not exist", id), err)
	} else if err != nil {
		fmt.Println("Other error:", err)
		return nil, err
	}

	return &ticket, nil
}
