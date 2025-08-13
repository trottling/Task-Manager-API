package mappers

import (
	"testing"
	"time"

	"server/internal/core/models"
)

func TestToTaskResponse(t *testing.T) {
	now := time.Now()
	task := models.Task{ID: 1, Name: "Test", Status: models.StatusNew, CreatedAt: now}
	resp := ToTaskResponse(task)

	if resp.ID != 1 || resp.Name != "Test" || resp.Status != string(models.StatusNew) {
		t.Errorf("mapper output incorrect: %+v", resp)
	}
}

func TestToListTasksResponse(t *testing.T) {
	tasks := []models.Task{
		{ID: 1, Name: "T1", Status: models.StatusNew},
		{ID: 2, Name: "T2", Status: models.StatusDone},
	}
	resp := ToListTasksResponse(tasks, 2)

	if len(resp.Items) != 2 || resp.Total != 2 {
		t.Errorf("expected total=2, got %+v", resp)
	}
}
