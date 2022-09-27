package controllers

import (
	"context"
	"database/sql"
	beans "grpc-demo-server/beans"
	daos "grpc-demo-server/daos"
	pb "grpc-demo-server/pb"
	"grpc-demo-server/utils/helpers"

	"github.com/volatiletech/null/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGpuModel(ctx context.Context, req *pb.CreateGpuModelRequest) (*pb.CreateGpuModelResponse, error) {
	bean := &beans.GpuModelsParams{
		VRam:             float64(req.GetVram()),
		OctaneBenchScore: int(req.GetOctaneBenchScore()),
		GpuNo:            int(req.GetGpuNo()),
		Available:        int(req.GetAvailable()),
		VramFree:         float64(req.GetVramFree()),
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

func (s *Server) UpdateGpuModel(ctx context.Context, req *pb.UpdateGpuModelRequest) (*pb.UpdateGpuModelResponse, error) {
	bean := &beans.UpdateGpuModelsParams{
		Id: req.Id,
		VRam: null.Float64{
			Float64: float64(req.GetVram()),
			Valid:   req.Vram != nil,
		},
		OctaneBenchScore: null.Int{
			Int:   int(req.GetOctaneBenchScore()),
			Valid: req.OctaneBenchScore != nil,
		},
		GpuNo: null.Int{
			Int:   int(req.GetGpuNo()),
			Valid: req.GpuNo != nil,
		},
		Available: null.Int{
			Int:   int(req.GetAvailable()),
			Valid: req.Available != nil,
		},
		VramFree: null.Float64{
			Float64: float64(req.GetVramFree()),
			Valid:   req.VramFree != nil,
		},
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

	updatedGpu, err := daos.UpdateGpu(ctx, s.DB, gpu, bean)
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateGpuModelResponse{
		Gpu: helpers.ConvertGpu(updatedGpu),
	}
	return resp, nil
}

func (s *Server) GetGpuModel(ctx context.Context, req *pb.GetGpuModelRequest) (*pb.GetGpuModelResponse, error) {
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

func (s *Server) GetGpuModels(ctx context.Context, req *pb.GetGpuModelsRequest) (*pb.GetGpuModelsResponse, error) {
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
