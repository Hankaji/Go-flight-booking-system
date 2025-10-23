// Package grpc
package grpc

import (
	"flight-booking-server/fbspb/plane"
	"flight-booking-server/grpc/services"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer() {
	lis, err := net.Listen("tcp4", "127.0.0.1:50051")
	if err != nil {
		panic(fmt.Sprintf("Err while listening %v", err))
	}

	s := grpc.NewServer()

	registerServices(s)
	reflection.Register(s)

	fmt.Println("GRPC listening on 127.0.0.1:50051")
	err = s.Serve(lis)
	if err != nil {
		panic(fmt.Sprintf("Err while serving %v", err))
	}
}

func registerServices(s *grpc.Server) {
	plane.RegisterPlaneServiceServer(s, &services.PlaneService{})
}
