package dto

// ErrorResponse описывает ошибку, возвращаемую клиенту
type ErrorResponse struct {
	Error string `json:"error"`
}
