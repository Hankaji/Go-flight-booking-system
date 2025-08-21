package routes

import (
	"flight-booking-server/controllers"
	"flight-booking-server/db"
	"flight-booking-server/repositories"
	"flight-booking-server/services"

	"github.com/gin-gonic/gin"
)

func RegisterFlightRoutes(r *gin.RouterGroup) {
	flightCotroller := controllers.FlightController{
		Service: services.FlightService{
			Repo: repositories.FlightRepo{
				DB: db.GetDBInstance(),
			},
		},
	}

	r.GET("flights", flightCotroller.GetFlights)
	r.GET("flights/:flightID", flightCotroller.GetFlightsByID)
	r.POST("flights", flightCotroller.CreateFlight)
}
