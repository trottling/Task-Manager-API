package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/internal/api/dto"
	"server/internal/core/logger"
	"server/internal/core/models"
	"server/internal/core/storage"
	"testing"
)

// Поднимаем чистый роутер на каждом тесте
func newTestServer() (*httptest.Server, *storage.Storage) {
	st := storage.NewStorage()
	log := logger.NewLogger("[test] ", 10)
	r := NewRouter(st, log)
	return httptest.NewServer(r), st
}

func mustJSON(t *testing.T, v any) *bytes.Reader {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return bytes.NewReader(b)
}

func decode[T any](t *testing.T, res *http.Response) T {
	t.Helper()
	defer func() {
		_ = res.Body.Close()
	}()

	var out T
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		t.Fatalf("decode: %v", err)
	}
	return out
}

func TestCreateTask_201_AndGetByID_200(t *testing.T) {
	srv, _ := newTestServer()
	defer srv.Close()

	// create
	req := dto.CreateTaskRequest{Name: "T1", Status: "new"}
	res, err := http.Post(srv.URL+"/tasks", "application/json", mustJSON(t, req))
	if err != nil {
		t.Fatalf("POST /tasks: %v", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", res.StatusCode)
	}
	created := decode[dto.TaskResponse](t, res)
	if created.Name != "T1" || created.Status != "new" {
		t.Fatalf("created mismatch: %+v", created)
	}

	// get by id
	res2, err := http.Get(srv.URL + "/tasks/1")
	if err != nil {
		t.Fatalf("GET /tasks/1: %v", err)
	}
	if res2.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res2.StatusCode)
	}
	got := decode[dto.TaskResponse](t, res2)
	if got.ID != 1 || got.Name != "T1" || got.Status != "new" {
		t.Fatalf("get mismatch: %+v", got)
	}
}

func TestCreateTask_400_BadJSON(t *testing.T) {
	srv, _ := newTestServer()
	defer srv.Close()

	res, err := http.Post(srv.URL+"/tasks", "application/json", bytes.NewReader([]byte("{bad json")))
	if err != nil {
		t.Fatalf("POST /tasks: %v", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.StatusCode)
	}
	_ = decode[dto.ErrorResponse](t, res)
}

func TestListTasks_ByStatus_200_AndPagination(t *testing.T) {
	srv, st := newTestServer()
	defer srv.Close()

	// seed
	_, _ = st.AddTask(&models.Task{Name: "A", Status: models.StatusNew})
	_, _ = st.AddTask(&models.Task{Name: "B", Status: models.StatusNew})
	_, _ = st.AddTask(&models.Task{Name: "C", Status: models.StatusDone})

	// list only "new", limit=1, offset=1 -> вернёт второй "new"
	res, err := http.Get(srv.URL + "/tasks?status=new&limit=1&offset=1")
	if err != nil {
		t.Fatalf("GET /tasks: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.StatusCode)
	}
	list := decode[dto.ListTasksResponse](t, res)
	if list.Total != 2 {
		t.Fatalf("expected total=2, got %d", list.Total)
	}
	if len(list.Items) != 1 || list.Items[0].Name != "B" {
		t.Fatalf("expected second 'new' -> B; got %+v", list.Items)
	}
}

func TestGetTaskByID_404(t *testing.T) {
	srv, _ := newTestServer()
	defer srv.Close()

	res, err := http.Get(srv.URL + "/tasks/999")
	if err != nil {
		t.Fatalf("GET /tasks/999: %v", err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", res.StatusCode)
	}
	_ = decode[dto.ErrorResponse](t, res)
}

func TestListTasks_400_InvalidStatus(t *testing.T) {
	srv, _ := newTestServer()
	defer srv.Close()

	res, err := http.Get(srv.URL + "/tasks?status=weird")
	if err != nil {
		t.Fatalf("GET /tasks: %v", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.StatusCode)
	}
	_ = decode[dto.ErrorResponse](t, res)
}
