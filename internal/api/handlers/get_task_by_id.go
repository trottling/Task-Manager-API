package handlers

import (
	"fmt"
	"net/http"
	"server/internal/api/dto"
	"server/internal/api/mappers"
	"server/internal/utils"
	"strconv"
)

// GetTaskByID GET /tasks/{id}
// Добавление задачи
func (rt *Router) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {

		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: fmt.Sprintf(ValidationError, "invalid task id")})
		return
	}

	task, err := rt.storage.GetByID(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, dto.ErrorResponse{Error: fmt.Sprintf(StorageError, err.Error())})
		return
	}

	utils.WriteJSON(w, http.StatusOK, mappers.ToTaskResponse(*task))
}
