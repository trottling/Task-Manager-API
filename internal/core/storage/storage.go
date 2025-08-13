package storage

import (
	"errors"
	"server/internal/core/models"
	"sync"
)

const (
	TaskNotFoundError = "task not found"
	EmptyStatusError  = "empty status"
)

// Storage примитивное in-memory хранилище для задач
type Storage struct {
	Tasks map[int]*models.Task
	Mutex sync.RWMutex
}

// NewStorage создаёт новое хранилище
func NewStorage() *Storage {
	return &Storage{
		Tasks: make(map[int]*models.Task),
	}
}

// AddTask добавляет задачу в хранилище
func (s *Storage) AddTask(task *models.Task) (int, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	taskID := len(s.Tasks) + 1
	s.Tasks[taskID] = task

	return taskID, nil
}

// GetByID ищет задачу по ID, если задача не найдена вернет ошибку TaskNotFoundError
func (s *Storage) GetByID(id int) (*models.Task, error) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	task, ok := s.Tasks[id]
	if !ok {
		return nil, errors.New(TaskNotFoundError)
	}
	return task, nil
}

// GetTasksByStatus ищет задачу по статусу, использует смещение и пагинацию
func (s *Storage) GetTasksByStatus(status string, limit, offset int) ([]models.Task, int, error) {
	if status == "" {
		return nil, 0, errors.New(EmptyStatusError)
	}

	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	var filtered []models.Task
	for _, task := range s.Tasks {
		if string(task.Status) == status {
			filtered = append(filtered, *task)
		}
	}

	total := len(filtered)

	// Пагинация
	if offset > total {
		return []models.Task{}, total, nil
	}

	end := offset + limit
	if end > total {
		end = total
	}

	return filtered[offset:end], total, nil
}
