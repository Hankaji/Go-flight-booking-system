package main

import (
	"flight-booking-server/db"
	"flight-booking-server/fbspb"
	"flight-booking-server/routes"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init DB
	dbConn := db.GetDBInstance()
	db.AutoMigrate(dbConn)

	go runGRPCServer()
	runRestServer()
}

func runRestServer() {
	router := gin.Default()

	routes.SetUpRoutes(router)

	router.Run(":8080")
}

type server struct {
	fbspb.FlightBookingSystemServiceServer
}

func runGRPCServer() {
	lis, err := net.Listen("tcp4", "127.0.0.1:50069")
	if err != nil {
		panic(fmt.Sprintf("Err while listening %v", err))
	}

	s := grpc.NewServer()

	fbspb.RegisterFlightBookingSystemServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		panic(fmt.Sprintf("Err while serving %v", err))
	}
}
