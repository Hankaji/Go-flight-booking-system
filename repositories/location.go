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

type ILocationRepo interface {
	IRepoRead[entities.Location]
}

type LocationRepo struct {
	DB *gorm.DB
}

func (repo *LocationRepo) GetAll() ([]entities.Location, error) {
	db := repo.DB

	var locations []entities.Location
	db.Find(&locations)

	return locations, nil
}

func (repo *LocationRepo) GetByID(id string) (*entities.Location, error) {
	db := repo.DB

	var location entities.Location
	err := db.First(&location, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("No location found with that ID")
		return nil, httperrors.NewHTTPErr(http.StatusNotFound, fmt.Errorf("location with ID %s does not exist", id), err)
	} else if err != nil {
		fmt.Println("Other error:", err)
		return nil, err
	}

	return &location, nil
}
