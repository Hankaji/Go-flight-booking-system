// Package services
package services

import (
	"context"
	"flight-booking-server/db"
	"flight-booking-server/fbspb/common"
	"flight-booking-server/fbspb/plane"
	"flight-booking-server/repositories"
	"flight-booking-server/services"
)

type PlaneService struct {
	plane.PlaneServiceServer
}

func (*PlaneService) GetAllPlanes(ctx context.Context, req *common.EmptyRequest) (*plane.PlaneResponse, error) {
	service := services.PlaneService{
		Repo: repositories.PlaneRepo{
			DB: db.GetDBInstance(),
		},
	}

	planes, err := service.GetAllPlanes()
	if err != nil {
		return nil, err
	}

	planeRes := make([]*plane.Plane, 0, len(planes))
	for _, dtoPlane := range planes {
		planeRes = append(planeRes, &plane.Plane{
			Id:         uint32(dtoPlane.ID),
			Model:      dtoPlane.Model,
			TotalSeats: uint32(dtoPlane.TotalSeats),
		},
		)
	}

	return &plane.PlaneResponse{
		Planes: planeRes,
	}, nil
}
