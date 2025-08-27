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

func (repo *FlightRepo) GetTickets(flightID string, filter *TicketFilter) ([]entities.Ticket, error) {
	db := repo.DB

	var tickets []entities.Ticket
	if err := db.Where("flight_id = ?", flightID).Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
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

func (repo *FlightRepo) CreateTicket(flightID uint, data dtos.CreateTicketRequest) (*entities.Ticket, error) {
	db := repo.DB

	// Check if flight exists
	var flight entities.Flight
	if err := db.First(&flight, flightID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, httperrors.NewHTTPErr(http.StatusNotFound, fmt.Errorf("flight with ID %d does not exist", flightID), err)
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
		FlightID: flightID,
		SeatID:   data.SeatID,
		Username: data.Username,
		Status:   entities.TicketApproved,
	}

	if err := db.Create(&newTicket).Error; err != nil {
		return nil, httperrors.NewHTTPErr(http.StatusInternalServerError, fmt.Errorf("could not create ticket for corresponding seat %d", data.SeatID), err)
	}

	return nil, nil
}

func (repo *FlightRepo) DeleteFlight(flightID string) error {
	db := repo.DB

	if err := db.Delete(&entities.Flight{}, flightID).Error; err != nil {
		return err
	}

	return nil
}

func (repo *FlightRepo) UpdateTicket(flightID, ticketID string) error {
	db := repo.DB

	if err := db.Delete(&entities.Flight{}, flightID).Error; err != nil {
		return err
	}

	return nil
}
