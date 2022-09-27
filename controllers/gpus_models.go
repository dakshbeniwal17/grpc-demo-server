package controllers

import (
	"context"
	"database/sql"
	beans "grpc-demo-server/beans"
	daos "grpc-demo-server/daos"
	pb "grpc-demo-server/pb"
	"grpc-demo-server/utils/helpers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGpu(ctx context.Context, req *pb.CreateGpuModelRequest) (*pb.CreateGpuModelResponse, error) {
	bean := &beans.GpuModelsParams{
		VRam:             float64(req.Vram),
		OctaneBenchScore: int(req.OctaneBenchScore),
	}

	gpu, err := daos.CreateGpu(ctx, s.DB, bean)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := &pb.CreateGpuModelResponse{
		Gpu: helpers.ConvertGpu(gpu),
	}
	return resp, nil
}

func (s *Server) UpdateGpu(ctx context.Context, req *pb.UpdateGpuModelRequest) (*pb.UpdateGpuModelResponse, error) {
	bean := &beans.GpuModelsParams{
		Id:               req.Id,
		VRam:             float64(req.GetVram()),
		OctaneBenchScore: int(req.GetOctaneBenchScore()),
	}

	if bean.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	gpu, err := daos.GetGpu(ctx, s.DB, bean.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "gpu not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find gpu")
	}

	gpu, err = daos.UpdateGpu(ctx, s.DB, gpu, bean)
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateGpuModelResponse{
		Gpu: helpers.ConvertGpu(gpu),
	}
	return resp, nil
}

func (s *Server) GetGpu(ctx context.Context, req *pb.GetGpuModelRequest) (*pb.GetGpuModelResponse, error) {
	id := req.Id
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Id is required")
	}

	gpu, err := daos.GetGpu(ctx, s.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "gpu not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find gpu")
	}

	resp := &pb.GetGpuModelResponse{
		Gpu: helpers.ConvertGpu(gpu),
	}
	return resp, nil
}

func (s *Server) GetGpus(ctx context.Context, req *pb.GetGpuModelsRequest) (*pb.GetGpuModelsResponse, error) {
	page := req.Page
	perPage := req.PerPage

	gpus, err := daos.GetGpus(ctx, s.DB, int(page), int(perPage))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find gpus")
	}

	respGpus := []*pb.GpuModels{}
	for _, gpu := range gpus {
		respGpus = append(respGpus, helpers.ConvertGpu(gpu))
	}
	resp := &pb.GetGpuModelsResponse{
		Gpus: respGpus,
	}
	return resp, nil
}
