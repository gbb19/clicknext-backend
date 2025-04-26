package dto

// ErrorResponse represents error response
type ErrorResponse struct {
	Error   string             `json:"error"`
	Details *map[string]string `json:"details"`
}
