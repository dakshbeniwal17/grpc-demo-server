package controllers

import (
	"database/sql"
	helloWorldPB "grpc-demo-server/pb"
)

type Server struct {
	helloWorldPB.UnimplementedGrpcDemoServer
	DB *sql.DB
}
