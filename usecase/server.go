package usecase

import (
	helloWorldPB "go-grpc/proto"
)

type Server struct {
	helloWorldPB.UnimplementedHelloWorldServer
}
