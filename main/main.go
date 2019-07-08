package main

import (
	"context"
	pb "head-start/protobuf"
	"log"
	"net"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) GetAllEmployee(ctx context.Context, req *pb.GetAllEmployeeRequest) (*pb.GetAllEmployeeResponse, error) {
	log.Println("Request received for GetAllEmployee")
	employeeResponse := &pb.GetAllEmployeeResponse{}
	var employeeList []*pb.Employee
	employeeList = append(employeeList, &pb.Employee{EmpName: "Name1", EmpEmail: "Email1"})
	employeeList = append(employeeList, &pb.Employee{EmpName: "Name2", EmpEmail: "Email2"})
	employeeList = append(employeeList, &pb.Employee{EmpName: "Name3", EmpEmail: "Email3"})
	employeeResponse.EmployeeList = employeeList
	return employeeResponse, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//s := grpc.NewServer()
	s := grpc.NewServer()

	grpcweb.WrapServer(s)
	var strarr []string
	strarr = append(strarr, "*")
	grpcweb.WithAllowedRequestHeaders(strarr)

	pb.RegisterEmployeeMgmtServiceServer(s, &server{})
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Serving Grpc service")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
