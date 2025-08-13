package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"server/internal/api/dto"
	"server/internal/api/mappers"
	"server/internal/utils"
)

// GetTaskByID - GET /tasks/{id}
func (rt *Router) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		rt.log.Info("bad id in path: %q", idStr)
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: fmt.Sprintf(ValidationErrorFmt, "invalid task id")})
		return
	}

	task, err := rt.storage.GetByID(id)
	if err != nil {
		// 404 - задача не найдена
		rt.log.Info("task not found id=%d: %v", id, err)
		utils.WriteJSON(w, http.StatusNotFound, dto.ErrorResponse{Error: fmt.Sprintf(StorageErrorFmt, err.Error())})
		return
	}

	rt.log.Info("task fetched id=%d", id)
	utils.WriteJSON(w, http.StatusOK, mappers.ToTaskResponse(*task))
}
