// Package controllers
package controllers

import (
	"errors"
	"flight-booking-server/controllers/queries"
	"flight-booking-server/dtos"
	"flight-booking-server/entities"
	httperrors "flight-booking-server/http-errors"
	"flight-booking-server/repositories"
	"flight-booking-server/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FlightController struct {
	Service services.FlightService
}

type FlightFilterQuery struct {
	PlaneID             *string                `form:"planeId"`
	Status              *entities.FlightStatus `form:"status"`
	TimeFrom            *time.Time             `form:"timeFrom"`
	TimeTo              *time.Time             `form:"timeTo"`
	DepartureLocationID *string                `form:"departureLocationId"`
	ArrivalLocationID   *string                `form:"arrivalLocationId"`
	Duration            *string                `form:"duration"`
	MinPrice            *string                `form:"minPrice"`
	MaxPrice            *string                `form:"maxPrice"`
}

func (q *FlightFilterQuery) toFilter() *repositories.FlightFilter {
	return &repositories.FlightFilter{
		PlaneID:             q.PlaneID,
		Status:              q.Status,
		TimeFrom:            q.TimeFrom,
		TimeTo:              q.TimeTo,
		DepartureLocationID: q.DepartureLocationID,
		ArrivalLocationID:   q.ArrivalLocationID,
		Duration:            q.Duration,
		MinPrice:            q.MinPrice,
		MaxPrice:            q.MaxPrice,
	}
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
		f.toFilter(),
	)
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

type FlightTicketFilterQuery struct {
	SeatID   *uint                  `form:"seatId"`
	Username *string                `form:"username"`
	Status   *entities.TicketStatus `form:"status"`
}

func (q *FlightTicketFilterQuery) toFilter() *repositories.FlightTicketFilter {
	return &repositories.FlightTicketFilter{
		SeatID:   q.SeatID,
		Username: q.Username,
		Status:   q.Status,
	}
}

func (con FlightController) GetTickets(ctx *gin.Context) {
	service := con.Service

	var q queries.PaginationQuery
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.Error(httperrors.NewHTTPErr(
			http.StatusBadRequest,
			errors.New("failed binding query"),
			err))
		return
	}

	var f FlightTicketFilterQuery
	if err := ctx.ShouldBindQuery(&f); err != nil {
		ctx.Error(httperrors.NewHTTPErr(
			http.StatusBadRequest,
			errors.New("failed binding query"),
			err))
		return
	}

	fmt.Println(f.toFilter())

	id := ctx.Param("flightID")

	tickets, err := service.GetAllTickets(id,
		repositories.PaginationFromQuery(q),
		f.toFilter())
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Tickets fetched successfully",
		"data":    tickets,
	})
}

func (con FlightController) GetTicketByID(ctx *gin.Context) {
	service := con.Service

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

	tickets, err := service.GetTicketByID(flightID, ticketID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Tickets fetched successfully",
		"data":    tickets,
	})
}
