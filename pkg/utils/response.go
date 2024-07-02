package utils

const StatusBadRequest = "Bad Request"
const StatusNotFound = "Not Found"
const StatusConflict = "Conflict Data"
const StatusOK = "OK"
const StatusUnauthorized = "Unauthorized"

const InvalidPageIndex = "Invalid Page Index"
const InvalidPageSize = "Invalid Page Size"
const MissingRequiredParams = "Missing required query parameters"

type Response struct {
	Header  HeaderDto   `json:"header"`
	Code    int         `json:"code,omitempty"`
	Status  string      `json:"status,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
