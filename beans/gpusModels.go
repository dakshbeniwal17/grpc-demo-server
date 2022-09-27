package beans

import (
	"github.com/volatiletech/null/v8"
)

type GpuModelsParams struct {
	VRam             float64 `json:"vram"`
	OctaneBenchScore int     `json:"octane_bench_score"`
	GpuNo            int     `json:"gpu_no"`
	Available        int     `json:"available"`
	VramFree         float64 `json:"vram_free"`
}

type UpdateGpuModelsParams struct {
	Id               string       `json:"id"`
	VRam             null.Float64 `json:"vram"`
	OctaneBenchScore null.Int     `json:"octane_bench_score"`
	GpuNo            null.Int     `json:"gpu_no"`
	Available        null.Int     `json:"available"`
	VramFree         null.Float64 `json:"vram_free"`
}
