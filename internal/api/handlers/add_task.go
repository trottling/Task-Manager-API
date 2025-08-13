package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/api/dto"
	"server/internal/api/mappers"
	"server/internal/api/validation"
	"server/internal/core/models"
	"server/internal/utils"
)

// CreateTask POST /tasks
// Добавление задачи
func (rt *Router) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: InvalidJSONError})
		return
	}

	if err := validation.ValidateCreateTaskRequest(req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	task := &models.Task{
		Name:   req.Name,
		Status: models.Status(req.Status),
	}

	id, err := rt.storage.AddTask(task)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, dto.ErrorResponse{Error: AddTaskError})
		return
	}

	rt.log.Info("created task id=%d", id)
	utils.WriteJSON(w, http.StatusCreated, mappers.ToTaskResponse(*task))
}
