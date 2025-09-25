package main

import (
	proto "ITUServer/grpc"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}

	client := proto.NewITUDatabaseClient(conn)

	students, err := client.GetStudents(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatalf("Not working")
	}

	for _, student := range students.Students {
		println(" - " + student)
	}
}
