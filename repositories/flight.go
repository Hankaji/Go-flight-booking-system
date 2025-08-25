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
	"gorm.io/gorm/clause"
)

type iFlightRepo interface {
	IRepoRead[entities.Flight]
	GetAllWithLocation() ([]entities.FlightDetailed, error)
}

type FlightRepo struct {
	DB *gorm.DB
}

func (repo *FlightRepo) GetAll() ([]entities.Flight, error) {
	db := repo.DB

	var flights []entities.Flight
	db.Find(&flights)

	return flights, nil
}

func (repo *FlightRepo) GetAllWithLocation() ([]entities.Flight, error) {
	db := repo.DB

	var flights []entities.Flight
	db.Preload(clause.Associations).Find(&flights)

	return flights, nil
}

func (repo *FlightRepo) GetByID(id string) (*entities.Flight, error) {
	db := repo.DB

	var flight entities.Flight
	if err := db.Preload(clause.Associations).First(&flight, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, httperrors.NewHTTPErr(http.StatusNotFound, fmt.Errorf("flight with ID %s does not exist", id), err)
	} else if err != nil {
		fmt.Println("Other error:", err)
		return nil, err
	}

	return &flight, nil
}

func (repo *FlightRepo) GetSeatAvailability(flightID string) ([]entities.SeatWithStatus, error) {
	db := repo.DB

	var seats []entities.SeatWithStatus
	if err := db.Raw(`
		SELECT s.*, 
			CASE WHEN t.id IS NULL THEN 'available' ELSE 'taken' END AS status
		FROM seats s
		JOIN flights f
			ON f.plane_id = s.plane_id
		LEFT JOIN tickets t 
			ON s.id = t.seat_id
			AND t.flight_id = ?
		WHERE f.id = ?
		`, flightID, flightID).Scan(&seats).Error; err != nil {
		return nil, err
	}

	return seats, nil
}

func (repo *FlightRepo) CreateFlight(data dtos.CreateFlightRequest) (*entities.Flight, error) {
	db := repo.DB

	newFlight := entities.Flight{
		DepartureTime:       data.DepartureTime,
		ArrivalTime:         data.ArrivalTime,
		DepartureLocationID: data.DepartureLocationID,
		ArrivalLocationID:   data.ArrivalLocationID,
		PlaneID:             data.Airplane,
		Price:               data.Price,
		Status:              entities.FlightScheduled,
	}

	db.Create(&newFlight)

	return nil, nil
}
