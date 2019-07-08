package main

import (
	"head-start/grpcservice"
	pb "head-start/protobuf"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmployeeMgmtServiceServer(s, &grpcservice.Server{})
	log.Println("Serving Grpc service")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
