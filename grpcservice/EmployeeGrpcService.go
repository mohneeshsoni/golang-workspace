package grpcservice

import (
	"context"
	"head-start/protobuf"
)

type Server struct {
}

func (server *Server) GetAllEmployee(ctx context.Context, req *protobuf.GetAllEmployeeRequest) (*protobuf.GetAllEmployeeResponse, error) {
	employeeResponse := &protobuf.GetAllEmployeeResponse{}
	var employeeList []*protobuf.Employee
	employeeList = append(employeeList, &protobuf.Employee{EmpName: "Name1", EmpEmail: "Email1"})
	employeeList = append(employeeList, &protobuf.Employee{EmpName: "Name2", EmpEmail: "Email2"})
	employeeList = append(employeeList, &protobuf.Employee{EmpName: "Name3", EmpEmail: "Email3"})
	return employeeResponse, nil
}
