package dto

//ProcResponse sfsd
type ProcResponse struct {
	Message    string `json:"message,omitempty"`
	Error      string `json:"error,omitempty"`
	IDProcess  int    `json:"idProcess,omitempty"`
	Data       int    `json:"data,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
}
