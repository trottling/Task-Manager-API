package storage

import (
	"testing"

	"server/internal/core/models"
)

func TestStorage_AddAndGet(t *testing.T) {
	st := NewStorage()

	task := &models.Task{Name: "Test", Status: models.StatusNew}
	id, err := st.AddTask(task)
	if err != nil {
		t.Fatalf("AddTask error: %v", err)
	}

	got, err := st.GetByID(id)
	if err != nil {
		t.Fatalf("GetByID error: %v", err)
	}
	if got.Name != "Test" {
		t.Errorf("expected name %q, got %q", "Test", got.Name)
	}
}

func TestStorage_GetTasksByStatus(t *testing.T) {
	st := NewStorage()
	_, _ = st.AddTask(&models.Task{Name: "T1", Status: models.StatusNew})
	_, _ = st.AddTask(&models.Task{Name: "T2", Status: models.StatusDone})

	tasks, total, err := st.GetTasksByStatus(string(models.StatusNew), 10, 0)
	if err != nil {
		t.Fatalf("GetTasksByStatus error: %v", err)
	}
	if total != 1 || tasks[0].Name != "T1" {
		t.Errorf("expected 1 task 'T1', got %v", tasks)
	}
}
