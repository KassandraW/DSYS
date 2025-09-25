package main

import (
	proto "ITUServer/grpc"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials())) //connects to server at localhost:5050. Insecure.newcredentials is used ti skip TLS encryption for simplification
	if err != nil {
		log.Fatalf("Not working")
	}

	client := proto.NewITUDatabaseClient(conn) //creates a client that can call RPC methods defined in the ITUDatabase service.

	students, err := client.GetStudents(context.Background(), &proto.Empty{}) //Client calls getstudents with an empty request and receives a students message containing the list of names.
	if err != nil {
		log.Fatalf("Not working")
	}

	for _, student := range students.Students { //prints each student returned from the request.
		println(" - " + student) 
	}
}
