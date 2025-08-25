// Package controllers
package controllers

import (
	"flight-booking-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeatController struct {
	Service services.SeatService
}

func (con SeatController) GetPlanes(ctx *gin.Context) {
	service := con.Service
	seats, err := service.GetAllSeats()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    seats,
	})
}

func (con SeatController) GetPlaneByID(ctx *gin.Context) {
	service := con.Service

	id := ctx.Param("seatID")

	seat, err := service.GetSeatByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    seat,
	})
}
