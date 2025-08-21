// Package routes provides HTTP route definitions and handlers
package routes

import (
	"flight-booking-server/middlewares"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.Use(cors.Default())
	api.Use(middlewares.ErrorHandler)

	api.GET("/health", healthCheck)

	RegisterFlightRoutes(api)
	RegisterLocationRoutes(api)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server is running",
	})
}
