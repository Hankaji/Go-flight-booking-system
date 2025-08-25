// Package routes
package routes

import (
	"flight-booking-server/controllers"
	"flight-booking-server/db"
	"flight-booking-server/repositories"
	"flight-booking-server/services"

	"github.com/gin-gonic/gin"
)

func RegisterTicketRoutes(r *gin.RouterGroup) {
	ticketController := controllers.TicketController{
		Service: services.TicketService{
			Repo: repositories.TicketRepo{
				DB: db.GetDBInstance(),
			},
		},
	}

	r.GET("tickets", ticketController.GetTickets)
	r.GET("tickets/:ticketID", ticketController.GetTicketByID)
	r.POST("tickets", ticketController.CreateTicket)
}
