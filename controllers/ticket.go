// Package controllers
package controllers

import (
	"flight-booking-server/dtos"
	httperrors "flight-booking-server/http-errors"
	"flight-booking-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	Service services.TicketService
}

func (con TicketController) GetTickets(ctx *gin.Context) {
	service := con.Service
	tickets, err := service.GetAllTickets()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    tickets,
	})
}

func (con TicketController) GetTicketByID(ctx *gin.Context) {
	service := con.Service

	id := ctx.Param("ticketID")

	ticket, err := service.GetTicketByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    ticket,
	})
}

func (con TicketController) CreateTicket(ctx *gin.Context) {
	var createTicketReq dtos.CreateTicketRequest
	if err := ctx.BindJSON(&createTicketReq); err != nil {
		ctx.Error(httperrors.NewHTTPErr(http.StatusBadRequest, nil, err))
		return
	}

	service := con.Service

	err := service.CreateTicket(createTicketReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Ticket created succesfully",
	})
}
