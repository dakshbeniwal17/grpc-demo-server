package controllers

import (
	"context"
	beans "grpc-demo-server/beans"
	daos "grpc-demo-server/daos"
	pb "grpc-demo-server/pb"
	"grpc-demo-server/utils/helpers"
)

func (s *Server) CreateHost(ctx context.Context, req *pb.CreateHostRequest) (*pb.CreateHostResponse, error) {
	bean := &beans.CreateHostParams{
		Name:    req.Name,
		Ram:     float64(req.Ram),
		Cores:   int(req.Cores),
		Threads: int(req.Threads),
		Ready:   int(req.Ready),
		JobUrl:  req.JobUrl,
		Token:   req.Token,
		Gpu:     req.Gpu,
	}
	// fmt.Printf("Host bean: %v\n", *bean)

	host, err := daos.Create(ctx, s.DB, bean)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateHostResponse{
		Host: helpers.ConvertHost(host),
	}
	return resp, nil
}
