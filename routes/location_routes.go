// Package routes
package routes

import (
	"flight-booking-server/controllers"
	"flight-booking-server/db"
	"flight-booking-server/repositories"
	"flight-booking-server/services"

	"github.com/gin-gonic/gin"
)

func RegisterLocationRoutes(r *gin.RouterGroup) {
	locationController := controllers.LocationController{
		Service: services.LocationService{
			Repo: repositories.LocationRepo{
				DB: db.GetDBInstance(),
			},
		},
	}

	r.GET("locations", locationController.GetLocations)
	r.GET("locations/:locationID", locationController.GetLocationByID)
}
