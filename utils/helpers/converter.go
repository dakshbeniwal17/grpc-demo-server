package helpers

import (
	"grpc-demo-server/pb"
	models "grpc-demo-server/repository"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertHost(host *models.Host) *pb.Host {
	return &pb.Host{
		Id:      host.ID,
		Name:    host.Name,
		Ram:     float32(host.RAM.Float64),
		Cores:   int64(host.Cores.Int),
		Threads: int64(host.Threads.Int),
		Updated: timestamppb.New(host.Updated.Time),
		Ready:   int64(host.Ready.Int),
		JobUrl:  host.JobURL.String,
		Token:   host.Token,
		Gpu:     host.Gpu.String,
	}
}

func ConvertGpu(gpu *models.GpuModel) *pb.GpuModels {
	return &pb.GpuModels{
		Id:               gpu.ID,
		Vram:             float32(gpu.Vram.Float64),
		OctaneBenchScore: int64(gpu.OctaneBenchScore.Int),
		GpuNo:            int64(gpu.GpuNo.Int),
		Available:        int64(gpu.Available.Int),
		VramFree:         float32(gpu.VramFree.Float64),
		Updated:          timestamppb.New(gpu.Updated.Time),
	}
}
