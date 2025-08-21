// Package controllers
package controllers

import (
	"flight-booking-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
	Service services.LocationService
}

func (con LocationController) GetLocations(ctx *gin.Context) {
	service := con.Service
	locations, err := service.GetAllLocations()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    locations,
	})
}

func (con LocationController) GetLocationByID(ctx *gin.Context) {
	service := con.Service

	id := ctx.Param("locationID")

	location, err := service.GetLocationByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    location,
	})
}
