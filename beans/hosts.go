package beans

type CreateHostParams struct {
	Name    string  `json:"name"`
	Ram     float64 `json:"ram"`
	Cores   int     `json:"cores"`
	Threads int     `json:"threads"`
	Ready   int     `json:"ready"`
	JobUrl  string  `json:"job_url"`
	Token   string  `json:"Token"`
	Gpu     string  `json:"gpu"`
}
