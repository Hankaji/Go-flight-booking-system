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

func (repo *FlightRepo) GetByID(id string) (*entities.FlightDetailed, error) {
	db := repo.DB

	var flight entities.Flight
	err := db.First(&flight, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("No flight found with that ID")
		return nil, httperrors.NewHTTPErr(http.StatusNotFound, fmt.Errorf("flight with ID %s does not exist", id), err)
	} else if err != nil {
		fmt.Println("Other error:", err)
		return nil, err
	}

	return nil, nil
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
