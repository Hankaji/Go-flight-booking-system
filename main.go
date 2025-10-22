package main

import (
	"context"
	"flight-booking-server/db"
	"flight-booking-server/fbspb/common"
	"flight-booking-server/fbspb/plane"
	"flight-booking-server/repositories"
	"flight-booking-server/routes"
	"flight-booking-server/services"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	plane.PlaneServiceServer
}

func runGRPCServer() {
	lis, err := net.Listen("tcp4", "127.0.0.1:50051")
	if err != nil {
		panic(fmt.Sprintf("Err while listening %v", err))
	}

	s := grpc.NewServer()

	plane.RegisterPlaneServiceServer(s, &server{})
	reflection.Register(s)

	fmt.Println("GRPC listening on 127.0.0.1:50051")
	err = s.Serve(lis)
	if err != nil {
		panic(fmt.Sprintf("Err while serving %v", err))
	}
}

func (server *server) GetAllPlanes(ctx context.Context, req *common.EmptyRequest) (*plane.PlaneResponse, error) {
	service := services.PlaneService{
		Repo: repositories.PlaneRepo{
			DB: db.GetDBInstance(),
		},
	}

	planes, err := service.GetAllPlanes()
	if err != nil {
		return nil, err
	}

	planeRes := make([]*plane.Plane, 0, len(planes))
	for _, dtoPlane := range planes {
		planeRes = append(planeRes, &plane.Plane{
			Id:         uint32(dtoPlane.ID),
			Model:      dtoPlane.Model,
			TotalSeats: uint32(dtoPlane.TotalSeats),
		},
		)
	}

	return &plane.PlaneResponse{
		Planes: planeRes,
	}, nil
}
