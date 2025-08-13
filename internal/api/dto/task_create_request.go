package dto

// CreateTaskRequest описывает данные, которые клиент отправляет для создания новой задачи
type CreateTaskRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Status      string  `json:"status"`
}
