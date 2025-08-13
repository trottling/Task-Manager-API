package validation

import (
	"errors"

	"server/internal/api/dto"
	"server/internal/core/models"
)

const (
	EmptyNameError   = "name must be non-empty"
	WrongStatusError = "status must be one of: \"new\", \"in_progress\", \"done\""
)

// IsValidStatus проверяет, что статус соответствует одному из допустимых значений
func IsValidStatus(s string) bool {
	return s == string(models.StatusNew) || s == string(models.StatusInProgress) || s == string(models.StatusDone)
}

// ValidateCreateTaskRequest проверяет корректность данных для создания новой задачи
func ValidateCreateTaskRequest(req dto.CreateTaskRequest) error {
	if req.Name == "" {
		return errors.New(EmptyNameError)
	}

	if !IsValidStatus(req.Status) {
		return errors.New(WrongStatusError)
	}

	return nil
}

// ProcessListQuery нормализует и валидирует параметры фильтрации/пагинации списка задач
func ProcessListQuery(q *dto.ListTasksQuery, defLimit, maxLimit int) error {

	// Проверяем статус
	if q.Status != nil && !IsValidStatus(*q.Status) {
		return errors.New(WrongStatusError)
	}

	// Если лимит меньше ноля или не указан, ставим дефолтный
	if q.Limit <= 0 {
		q.Limit = defLimit
	}

	// Если лимит больше максимального, ставим максимальный
	if q.Limit > maxLimit {
		q.Limit = maxLimit
	}

	// Если смещение меньше ноля или не указан, ставим дефолтный
	if q.Offset < 0 {
		q.Offset = 0
	}

	return nil
}
