package main

import (
	"log"
	"net"

	helloWorldPB "go-grpc/proto"
	usecase "go-grpc/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := ":9000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	helloWorldPB.RegisterHelloWorldServer(grpcServer, &usecase.Server{})

	log.Printf("GRPC server listening on %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
