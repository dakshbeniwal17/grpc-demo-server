package usecase

import (
	helloWorldPB "grpc-demo-server/proto"
)

type Server struct {
	helloWorldPB.UnimplementedHelloWorldServer
}
