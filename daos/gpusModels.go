package daos

import (
	"context"
	"log"

	"grpc-demo-server/beans"
	models "grpc-demo-server/repository"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func CreateGpu(ctx context.Context, exec boil.ContextExecutor, bean *beans.GpuModelsParams) (*models.GpuModel, error) {
	gpu := &models.GpuModel{
		Vram:             null.Float64From(bean.VRam),
		OctaneBenchScore: null.IntFrom(bean.OctaneBenchScore),
		GpuNo:            null.IntFrom(bean.GpuNo),
		SlotNo:           null.IntFrom(bean.SlotNo),
		Available:        null.IntFrom(bean.Available),
		VramFree:         null.Float64From(bean.VramFree),
	}
	if err := gpu.Insert(ctx, exec, boil.Infer()); err != nil {
		log.Printf("error while creating gpu: %v", gpu)
		return nil, err
	}
	return gpu, nil
}

func GetGpu(ctx context.Context, exec boil.ContextExecutor, id string) (*models.GpuModel, error) {
	gpu, err := models.GpuModels(models.GpuModelWhere.ID.EQ(id)).One(ctx, exec)
	if err != nil {
		log.Printf("error while fetching gpu where id = %s", id)
		return nil, err
	}
	return gpu, nil
}

func GetGpus(ctx context.Context, exec boil.ContextExecutor, page int, perPage int) (models.GpuModelSlice, error) {
	if perPage == 0 {
		perPage = 10
	}
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * perPage
	gpus, err := models.GpuModels(
		qm.Limit(perPage),
		qm.Offset(offset),
	).All(ctx, exec)
	if err != nil {
		log.Printf("error while fetching gpus where page = %d and per_page = %d", page, perPage)
		return nil, err
	}
	return gpus, nil
}

func UpdateGpu(ctx context.Context, exec boil.ContextExecutor, gpu *models.GpuModel, bean *beans.UpdateGpuModelsParams) (*models.GpuModel, error) {
	if bean.VRam.Valid {
		gpu.Vram = bean.VRam
	}
	if bean.OctaneBenchScore.Valid {
		gpu.OctaneBenchScore = bean.OctaneBenchScore
	}
	if bean.GpuNo.Valid {
		gpu.GpuNo = bean.GpuNo
	}
	if bean.SlotNo.Valid {
		gpu.SlotNo = bean.SlotNo
	}
	if bean.VramFree.Valid {
		gpu.VramFree = bean.VramFree
	}
	if bean.Available.Valid {
		gpu.Available = bean.Available
	}
	if _, err := gpu.Update(ctx, exec, boil.Infer()); err != nil {
		log.Printf("error while updating gpu: %v", gpu)
		return nil, err
	}
	return gpu, nil
}
