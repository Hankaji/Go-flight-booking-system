// Package controllers
package controllers

import (
	"flight-booking-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlaneController struct {
	Service services.PlaneService
}

func (con PlaneController) GetPlanes(ctx *gin.Context) {
	service := con.Service
	locations, err := service.GetAllPlanes()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    locations,
	})
}

func (con PlaneController) GetPlaneByID(ctx *gin.Context) {
	service := con.Service

	id := ctx.Param("locationID")

	location, err := service.GetPlanByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successfully",
		"data":    location,
	})
}
