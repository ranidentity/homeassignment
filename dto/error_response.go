package dto

type ErrorDetail interface{}
type ErrorResponse struct {
	Error   string      `json:"error"`
	Details ErrorDetail `json:"details"`
}
