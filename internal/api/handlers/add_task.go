package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"server/internal/api/dto"
	"server/internal/api/mappers"
	"server/internal/api/validation"
	"server/internal/core/models"
	"server/internal/utils"
)

// CreateTask - POST /tasks
func (rt *Router) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rt.log.Info("json decode error: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: fmt.Sprintf(InvalidJSONErrorFmt, err.Error())})
		return
	}

	if err := validation.ValidateCreateTaskRequest(req); err != nil {
		rt.log.Info("validation failed: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: fmt.Sprintf(ValidationErrorFmt, err.Error())})
		return
	}

	task := &models.Task{
		Name:   req.Name,
		Status: models.Status(req.Status),
	}

	id, err := rt.storage.AddTask(task)
	if err != nil {
		rt.log.Error("storage add error: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, dto.ErrorResponse{Error: fmt.Sprintf(StorageErrorFmt, err.Error())})
		return
	}

	rt.log.Info("task created id=%d name=%q status=%s", id, task.Name, task.Status)
	utils.WriteJSON(w, http.StatusCreated, mappers.ToTaskResponse(*task))
}
