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

type ISeatRepo interface {
	IRepoRead[entities.Location]
	GetSeatAvailability(flightID string) ([]entities.SeatWithStatus, error)
}

type SeatRepo struct {
	DB *gorm.DB
}

func (repo *SeatRepo) GetAll() ([]entities.Seat, error) {
	db := repo.DB

	var seats []entities.Seat
	if err := db.Find(&seats).Error; err != nil {
		return nil, err
	}

	return seats, nil
}

func (repo *SeatRepo) GetByID(id string) (*entities.Seat, error) {
	db := repo.DB

	var seat entities.Seat
	if err := db.First(&seat, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, httperrors.NewHTTPErr(http.StatusNotFound, fmt.Errorf("seat with ID %s does not exist", id), err)
	} else if err != nil {
		fmt.Println("Other error:", err)
		return nil, err
	}

	return &seat, nil
}
