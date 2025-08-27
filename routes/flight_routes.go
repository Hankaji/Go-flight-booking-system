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

	r.GET("flights", controller.GetFlights)
	r.GET("flights/:flightID", controller.GetFlightsByID)
	r.GET("flights/:flightID/seats", controller.GetSeatAvailability)
	r.GET("flights/:flightID/tickets", controller.GetTickets)

	r.POST("flights", controller.CreateFlight)

	r.DELETE("flights/:flightID", controller.DeleteFlight)
}
