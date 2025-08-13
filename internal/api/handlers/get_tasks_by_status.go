package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"server/internal/api/dto"
	"server/internal/api/mappers"
	"server/internal/api/validation"
	"server/internal/utils"
)

// ListTasks - GET /tasks?status=...&limit=...&offset=...
func (rt *Router) ListTasks(w http.ResponseWriter, r *http.Request) {
	var q dto.ListTasksQuery

	if s := r.URL.Query().Get("status"); s != "" {
		q.Status = &s
	}

	if lim, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		q.Limit = lim
	}

	if off, err := strconv.Atoi(r.URL.Query().Get("offset")); err == nil {
		q.Offset = off
	}

	if err := validation.ProcessListQuery(&q, ListQueryDefaultLimit, ListQueryMaxLimit); err != nil {
		rt.log.Info("list query validation failed: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: fmt.Sprintf(ValidationErrorFmt, err.Error())})
		return
	}

	tasks, total, err := rt.storage.GetTasksByStatus(*q.Status, q.Limit, q.Offset)
	if err != nil {
		rt.log.Error("storage list error: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, dto.ErrorResponse{Error: fmt.Sprintf(StorageErrorFmt, err.Error())})
		return
	}

	rt.log.Info("tasks listed status=%s limit=%d offset=%d total=%d", *q.Status, q.Limit, q.Offset, total)
	utils.WriteJSON(w, http.StatusOK, mappers.ToListTasksResponse(tasks, total))
}
