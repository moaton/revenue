package dto

type Chart struct {
	Data []int64 `json:"data"`
}

type GetChartsResponse struct {
	Data []int64 `json:"data"`
}
