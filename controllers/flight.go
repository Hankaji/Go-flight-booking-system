// Package controllers
package controllers

import (
	"errors"
	"flight-booking-server/controllers/queries"
	"flight-booking-server/dtos"
	httperrors "flight-booking-server/http-errors"
	"flight-booking-server/repositories"
	"flight-booking-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlightController struct {
	Service services.FlightService
}

func (con FlightController) GetFlights(ctx *gin.Context) {
	var q queries.PaginationQuery
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.Error(httperrors.NewHTTPErr(
			http.StatusBadRequest,
			errors.New("failed binding query"),
			err))
		return
	}

	flights, err := con.Service.GetAllFlights(
		repositories.PaginationFromQuery(q),
		nil)
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

func (con FlightController) GetSeatAvailability(ctx *gin.Context) {
	service := con.Service

	id := ctx.Param("flightID")

	seat, err := service.GetAllSeatsAvailability(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    seat,
	})
}
