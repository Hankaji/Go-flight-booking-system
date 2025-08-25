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

type FlightRepo struct {
	DB *gorm.DB
}

type FlightFilter struct {
	PlaneID *string
}

func (repo *FlightRepo) GetAllWithParams(pagination *Pagination, filter *FlightFilter) ([]entities.Flight, error) {
	db := repo.DB.Model(&entities.Flight{})

	var flights []entities.Flight

	if pagination == nil {
		pagination = DefaultPagination()
	}
	db.Limit(int(pagination.limit))
	db.Offset(int(pagination.index))

	if filter != nil {
		if filter.PlaneID != nil {
			db.Where("plane_id = ?", *filter.PlaneID)
		}
	}

	db.Find(&flights)

	return flights, nil
}

func (repo *FlightRepo) GetAll() ([]entities.Flight, error) {
	return repo.GetAllWithParams(nil, nil)
}

func (repo *FlightRepo) GetAllWithLocationWithParams(pagination *Pagination, filter *FlightFilter) ([]entities.Flight, error) {
	db := repo.DB.Model(&entities.Flight{})

	var flights []entities.Flight

	if pagination == nil {
		pagination = DefaultPagination()
	}

	db.Offset(int((pagination.index - 1) * pagination.limit))
	db.Limit(int(pagination.limit))

	if filter != nil {
		if filter.PlaneID != nil {
			db.Where("plane_id = ?", *filter.PlaneID)
		}
	}

	db.Preload(clause.Associations).Find(&flights)

	return flights, nil
}

func (repo *FlightRepo) GetAllWithLocation() ([]entities.Flight, error) {
	return repo.GetAllWithLocationWithParams(nil, nil)
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
