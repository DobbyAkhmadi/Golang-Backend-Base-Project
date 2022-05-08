package entity

type ErrorResponse struct {
	Status   bool        `json:"status"`
	Code     int         `json:"code,omitempty"`
	Title    string      `json:"title,omitempty"`
	Message  string      `json:"message,omitempty"`
	Errors   interface{} `json:"errors,omitempty"`
	TraceID  string      `json:"traceId,omitempty"`
	Instance string      `json:"instance,omitempty"`
}

type ErrValidation struct {
	Tag    string      `json:"tag,omitempty"`
	Value  interface{} `json:"value,omitempty"`
	Reason string      `json:"reason,omitempty"`
}
