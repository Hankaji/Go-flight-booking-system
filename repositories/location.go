// Package repositories
package repositories

import (
	"flight-booking-server/entities"

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
	db.First(&location, id)

	return &location, nil
}
