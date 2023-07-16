package utils

type Response struct {
	Header HeaderDto   `json:"header"`
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
