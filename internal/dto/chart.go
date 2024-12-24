package dto

type ChartJsonPoint struct {
	V []float64 `json:"v"`
	C []float64 `json:"c"`
}

type ChartJson struct {
	Data   ChartJsonData   `json:"data"`
	Status ChartJsonStatus `json:"status"`
}

type ChartJsonData struct {
	Points map[string]ChartJsonPoint `json:"points"`
}

type ChartJsonStatus struct {
	ErrorCode string `json:"error_code"`
}
