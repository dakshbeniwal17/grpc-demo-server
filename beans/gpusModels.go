package beans

type GpuModelsParams struct {
	Id               string  `json:"id"`
	VRam             float64 `json:"vram"`
	OctaneBenchScore int     `json:"octane_bench_score"`
}
