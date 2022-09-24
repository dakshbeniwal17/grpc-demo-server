package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"grpc-demo-server/db"
	helloWorldPB "grpc-demo-server/proto"
	usecase "grpc-demo-server/usecase"
	"grpc-demo-server/utils/helpers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := ":9000"

	config, err := helpers.LoadEnvFile()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}

	database, err := db.PostgresConnect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] initializing postgres db: %+v\n", err)
		return
	}

	db.RunDBMigration(config.Get("db.migrationUrl").(string), database)

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
