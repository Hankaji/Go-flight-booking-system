// Package controllers
package controllers

import (
	"flight-booking-server/dtos"
	"flight-booking-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlightController struct {
	Service services.FlightService
}

func (con FlightController) GetFlights(ctx *gin.Context) {
	flights, err := con.Service.GetAllFlights()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    flights,
	})
}

func (con FlightController) GetFlightsByID(ctx *gin.Context) {
	id := ctx.Param("flightID")

	flight, err := con.Service.GetFlightByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    flight,
	})
}

func (con FlightController) CreateFlight(ctx *gin.Context) {
	var createFlightReq dtos.CreateFlightRequest
	if err := ctx.BindJSON(&createFlightReq); err != nil {
		ctx.Error(err)
		return
	}

	err := con.Service.CreateFlight(createFlightReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Flight created succesfully",
	})
}
