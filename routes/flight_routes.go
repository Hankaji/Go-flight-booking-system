package routes

import (
	"flight-booking-server/controllers"
	"flight-booking-server/db"
	"flight-booking-server/repositories"
	"flight-booking-server/services"

	"github.com/gin-gonic/gin"
)

func RegisterFlightRoutes(r *gin.RouterGroup) {
	controller := controllers.FlightController{
		Service: services.FlightService{
			Repo: repositories.FlightRepo{
				DB: db.GetDBInstance(),
			},
		},
	}

	// Flights
	r.GET("flights", controller.GetFlights)
	r.GET("flights/:flightID", controller.GetFlightsByID)
	r.POST("flights", controller.CreateFlight)
	r.DELETE("flights/:flightID", controller.DeleteFlight)

	// Seats
	r.GET("flights/:flightID/seats", controller.GetSeatAvailability)

	// Tickets
	r.GET("flights/:flightID/tickets", controller.GetTickets)
	r.POST("flights/:flightID/tickets", controller.CreateTicket)
	r.PATCH("flights/:flightID/tickets/:ticketID", controller.UpdateTicket)
}
