package utils

type Response struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
