package main

import (
	"flight-booking-server/db"
	"flight-booking-server/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init DB
	dbConn := db.GetDBInstance()
	db.AutoMigrate(dbConn)

	router := gin.Default()

	routes.SetUpRoutes(router)

	router.Run(":8080")
}
