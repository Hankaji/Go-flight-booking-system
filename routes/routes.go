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
	// api.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"}, // or "*" for dev
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	api.Use(middlewares.ErrorHandler)

	api.GET("/health", healthCheck)

	RegisterFlightRoutes(api)
	RegisterLocationRoutes(api)
	RegisterPlaneRoutes(api)
	RegisterTicketRoutes(api)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server is running",
	})
}
