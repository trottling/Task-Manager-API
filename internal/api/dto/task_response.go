package dto

import "time"

// TaskResponse формат задачи, который возвращается клиенту через API
type TaskResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// ListTasksResponse используется для возврата коллекции задач
type ListTasksResponse struct {
	Items []TaskResponse `json:"items"`
	Total int            `json:"total"`
}
