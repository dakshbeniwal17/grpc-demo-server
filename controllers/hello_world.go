package controllers

import (
	"context"
	helloWorldPB "grpc-demo-server/pb"
)

func (s *Server) GetMessage(ctx context.Context, in *helloWorldPB.GetMessageRequest) (*helloWorldPB.GetMessageResponse, error) {
	return &helloWorldPB.GetMessageResponse{
		Message: "Hello World!",
	}, nil
}
