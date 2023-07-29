package utils

var StatusBadRequest = "Bad Request"
var StatusNotFound = "Not Found"
var StatusConflict = "Conflict Data"
var StatusOK = "OK"
var StatusUnauthorized = "Unauthorized"
var StatusForbidden = "Forbidden"

type Response struct {
	Header HeaderDto   `json:"header"`
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
