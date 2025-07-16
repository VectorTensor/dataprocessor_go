package main

import (
	"fmt"
	"log"
	"main/codegen"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"main/services"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {

		log.Fatalf("Failed to listen: %v", err)

	}

	s := grpc.NewServer()
	reflection.Register(s)

	codegen.RegisterDataprocessorServer(s, &services.DataProcessorServer{})

	fmt.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {

		log.Fatalf("Failed to server: %v", err)

	}

}
