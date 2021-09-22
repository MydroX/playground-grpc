package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/MydroX/geotwin-grpctest/internal/positions"
	pbPositions "github.com/MydroX/geotwin-grpctest/internal/positions/proto"
)

func GRPCServer() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	s := grpc.NewServer()

	positionsServer := positions.NewPositionsServer()

	pbPositions.RegisterPositionServiceServer(s, positionsServer)

	log.Println("Starting GRPC server...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
