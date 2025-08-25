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
	if err := db.First(&plane, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, httperrors.NewHTTPErr(http.StatusNotFound, fmt.Errorf("plane with ID %s does not exist", id), err)
	} else if err != nil {
		fmt.Println("Other error:", err)
		return nil, err
	}

	return &plane, nil
}
