// Package controllers
package controllers

import (
	"errors"
	"flight-booking-server/controllers/queries"
	"flight-booking-server/dtos"
	httperrors "flight-booking-server/http-errors"
	"flight-booking-server/repositories"
	"flight-booking-server/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FlightController struct {
	Service services.FlightService
}

type FlightFilterQuery struct {
	PlaneID *string `form:"plane_id"`
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

	if err := ValidatePagination(q); err != nil {
		ctx.Error(*err)
		return
	}

	var f FlightFilterQuery
	if err := ctx.ShouldBindQuery(&f); err != nil {
		ctx.Error(httperrors.NewHTTPErr(
			http.StatusBadRequest,
			errors.New("failed binding query"),
			err))
		return
	}

	fmt.Println("Filter: ", f)

	flights, err := con.Service.GetAllFlights(
		repositories.PaginationFromQuery(q),
		&repositories.FlightFilter{
			PlaneID: f.PlaneID,
		})
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
	idParam := ctx.Param("flightID")
	id, parseErr := strconv.ParseUint(idParam, 10, 32)
	if parseErr != nil {
		ctx.Error(httperrors.NewHTTPErr(http.StatusBadRequest, errors.New("invalid flightID"), parseErr))
		return
	}

	flight, err := con.Service.GetFlightByID(uint(id))
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

func (con FlightController) CreateTicket(ctx *gin.Context) {
	var createTicketReq dtos.CreateTicketRequest
	if err := ctx.BindJSON(&createTicketReq); err != nil {
		ctx.Error(httperrors.NewHTTPErr(http.StatusBadRequest,
			nil,
			err))
		return
	}

	idParam := ctx.Param("flightID")
	id, parseErr := strconv.ParseUint(idParam, 10, 32)
	if parseErr != nil {
		ctx.Error(httperrors.NewHTTPErr(http.StatusBadRequest, errors.New("invalid flightID"), parseErr))
		return
	}

	service := con.Service

	err := service.CreateTicket(uint(id), createTicketReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Flight ticket created succesfully",
	})
}

func (con FlightController) UpdateTicket(ctx *gin.Context) {
	var updateTicketReq dtos.UpdateTicketRequest
	if err := ctx.BindJSON(&updateTicketReq); err != nil {
		ctx.Error(httperrors.NewHTTPErr(http.StatusBadRequest,
			nil,
			err))
		return
	}

	var flightID uint
	{
		idParam := ctx.Param("flightID")
		id, parseErr := strconv.ParseUint(idParam, 10, 32)
		if parseErr != nil {
			ctx.Error(httperrors.NewHTTPErr(http.StatusBadRequest, errors.New("invalid flightID"), parseErr))
			return
		}

		flightID = uint(id)
	}

	var ticketID uint
	{
		idParam := ctx.Param("ticketID")
		id, parseErr := strconv.ParseUint(idParam, 10, 32)
		if parseErr != nil {
			ctx.Error(httperrors.NewHTTPErr(http.StatusBadRequest, errors.New("invalid ticketID"), parseErr))
			return
		}

		ticketID = uint(id)
	}

	service := con.Service

	err := service.UpdateTicket(flightID, ticketID, updateTicketReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Flight ticket updated succesfully",
	})
}

func (con FlightController) UpdateFlight(ctx *gin.Context) {
	var updateTicketReq dtos.UpdateFlightRequest
	if err := ctx.BindJSON(&updateTicketReq); err != nil {
		ctx.Error(httperrors.NewHTTPErr(http.StatusBadRequest,
			nil,
			err))
		return
	}

	var flightID uint
	{
		idParam := ctx.Param("flightID")
		id, parseErr := strconv.ParseUint(idParam, 10, 32)
		if parseErr != nil {
			ctx.Error(httperrors.NewHTTPErr(http.StatusBadRequest, errors.New("invalid flightID"), parseErr))
			return
		}

		flightID = uint(id)
	}

	service := con.Service

	err := service.UpdateFlight(flightID, updateTicketReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Flight updated succesfully",
	})
}

func (con FlightController) DeleteFlight(ctx *gin.Context) {
	service := con.Service

	id := ctx.Param("flightID")

	err := service.DeleteFlight(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Flight deleted succesfully",
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

func (con FlightController) GetTickets(ctx *gin.Context) {
	service := con.Service

	id := ctx.Param("flightID")

	tickets, err := service.GetAllTickets(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Tickets fetched successfully",
		"data":    tickets,
	})
}
