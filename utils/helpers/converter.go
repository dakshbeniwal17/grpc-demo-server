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
