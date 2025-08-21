// Package repositories
package repositories

import (
	"flight-booking-server/dtos"
	"flight-booking-server/entities"
	"fmt"

	"gorm.io/gorm"
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

func (repo *FlightRepo) GetAllWithLocation() ([]entities.FlightDetailed, error) {
	db := repo.DB

	var flights []entities.FlightDetailed
	db.Find(&flights)

	fmt.Println("FlightRepo: ", flights)

	return flights, nil
}

func (repo *FlightRepo) GetByID(id string) (*entities.FlightDetailed, error) {
	db := repo.DB

	var flight entities.Flight
	db.First(&flight, id)

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
