// Package repositories
package repositories

import (
	"errors"
	"flight-booking-server/dtos"
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

func (repo *TicketRepo) Create(data dtos.CreateTicketRequest) (*entities.Ticket, error) {
	db := repo.DB

	// Check if flight exists
	var flight entities.Seat
	if err := db.First(&flight, data.FlightID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, httperrors.NewHTTPErr(http.StatusNotFound, fmt.Errorf("flight with ID %d does not exist", data.SeatID), err)
		}
		return nil, err
	}

	// Check if seat exists
	var seat entities.Seat
	if err := db.First(&seat, data.SeatID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, httperrors.NewHTTPErr(http.StatusNotFound, fmt.Errorf("seat with ID %d does not exist", data.SeatID), err)
		}
		return nil, err
	}

	// Check if seat is booked or not
	var existingTicket entities.Ticket
	if err := db.Where("seat_id = ? AND status != ?", data.SeatID, entities.TicketCancelled).
		First(&existingTicket).Error; err == nil {
		return nil, httperrors.NewHTTPErr(http.StatusConflict, fmt.Errorf("seat %d is already booked", data.SeatID), err)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	newTicket := entities.Ticket{
		SeatID:   data.SeatID,
		Username: data.Username,
		Status:   entities.TicketApproved,
	}

	if err := db.Create(&newTicket).Error; err != nil {
		return nil, httperrors.NewHTTPErr(http.StatusInternalServerError, fmt.Errorf("could not create ticket for corresponding seat %d", data.SeatID), err)
	}

	return nil, nil
}
