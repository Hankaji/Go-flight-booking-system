// Package routes
package routes

import (
	"flight-booking-server/controllers"
	"flight-booking-server/db"
	"flight-booking-server/repositories"
	"flight-booking-server/services"

	"github.com/gin-gonic/gin"
)

func RegisterPlaneRoutes(r *gin.RouterGroup) {
	controller := controllers.PlaneController{
		Service: services.PlaneService{
			Repo: repositories.PlaneRepo{
				DB: db.GetDBInstance(),
			},
		},
	}

	r.GET("planes", controller.GetPlanes)
	r.GET("planes/:planeID", controller.GetPlaneByID)
}
