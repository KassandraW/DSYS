package main

import (
	proto "ITUServer/grpc"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ITU_databaseServer struct { //implements the gRPC service efinted in the .proto file.
	proto.UnimplementedITUDatabaseServer //Embeds UnimplementedITUDatabaseServer to satisfy the interface
	students []string // Includes a slice of students.
}

func (s *ITU_databaseServer) GetStudents(ctx context.Context, in *proto.Empty) (*proto.Students, error) { //takes a context and an empty request
	return &proto.Students{Students: s.students}, nil //returns a Students message containing the list of student names and a possible error.
}

func main() { //initializes server
	server := &ITU_databaseServer{students: []string{}}
	server.students = append(server.students, "John")
	server.students = append(server.students, "Jane")
	server.students = append(server.students, "Alice")
	server.students = append(server.students, "Bob")

	server.start_server() //starts the gRPC server
}

func (s *ITU_databaseServer) start_server() {
	grpcServer := grpc.NewServer() //Creates a new gRPC server instance
	listener, err := net.Listen("tcp", ":5050") //Listens on TCP port 5050 using net.Listen
	if err != nil {
		log.Fatalf("Did not work")
	}

	proto.RegisterITUDatabaseServer(grpcServer, s) //registers the server implementation with gRPC

	err = grpcServer.Serve(listener) // starts serving incoming requests

	if err != nil {
		log.Fatalf("Did not work")
	}

}
