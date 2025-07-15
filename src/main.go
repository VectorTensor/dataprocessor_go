package main

import (
	"context"
	"fmt"
	"log"
	"main/codegen"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type greeterServer struct {
	codegen.UnimplementedDataprocessorServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *codegen.HelloRequest) (*codegen.HelloReply, error) {

	log.Printf("Receieved %s", req.GetName())

	return &codegen.HelloReply{
		Message: "Hello " + req.GetName(),
	}, nil

}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {

		log.Fatalf("Failed to listen: %v", err)

	}

	s := grpc.NewServer()
	reflection.Register(s)

	codegen.RegisterDataprocessorServer(s, &greeterServer{})

	fmt.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {

		log.Fatalf("Failed to server: %v", err)

	}

}
