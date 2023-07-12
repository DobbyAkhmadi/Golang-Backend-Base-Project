package utils

type ErrorResponse struct {
	Code     int         `json:"code,omitempty"`
	Status   string      `json:"status"`
	Errors   interface{} `json:"errors,omitempty"`
	TraceID  string      `json:"traceId,omitempty"`
	Instance string      `json:"instance,omitempty"`
}

type ErrValidation struct {
	Tag    string      `json:"tag,omitempty"`
	Value  interface{} `json:"value,omitempty"`
	Reason string      `json:"reason,omitempty"`
}
