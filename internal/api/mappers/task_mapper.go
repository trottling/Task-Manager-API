package mappers

import (
	"server/internal/api/dto"
	"server/internal/core/models"
)

// ToTaskResponse преобразует доменную модель Task в DTO для ответа API
func ToTaskResponse(t models.Task) dto.TaskResponse {
	return dto.TaskResponse{
		ID:        t.ID,
		Name:      t.Name,
		Status:    string(t.Status),
		CreatedAt: t.CreatedAt,
	}
}

// ToListTasksResponse преобразует список доменных моделей Task в DTO-структуру списка
func ToListTasksResponse(items []models.Task, total int) dto.ListTasksResponse {
	out := make([]dto.TaskResponse, 0, len(items))
	for _, it := range items {
		out = append(out, ToTaskResponse(it))
	}
	return dto.ListTasksResponse{
		Items: out,
		Total: total,
	}
}
