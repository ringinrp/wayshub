package dto

type SuccessResult struct {
	Code   string      `json:"code"`
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
