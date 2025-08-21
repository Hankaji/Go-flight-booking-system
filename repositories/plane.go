// Package repositories
package repositories

import (
	"flight-booking-server/entities"

	"gorm.io/gorm"
)

type IPlaneRepo interface {
	IRepoRead[entities.Location]
}

type PlaneRepo struct {
	DB *gorm.DB
}

func (repo *PlaneRepo) GetAll() ([]entities.Plane, error) {
	db := repo.DB

	var planes []entities.Plane
	db.Find(&planes)

	return planes, nil
}

func (repo *PlaneRepo) GetByID(id string) (*entities.Plane, error) {
	db := repo.DB

	var plane entities.Plane
	db.First(&plane, id)

	return &plane, nil
}
