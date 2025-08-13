package handlers

import (
	"net/http"

	"server/internal/core/logger"
	"server/internal/core/storage"
)

const (
	InvalidJSONError      = "InvalidJSON"
	AddTaskError          = "AddTaskError"
	ListQueryMaxLimit     = 200
	ListQueryDefaultLimit = 50
)

type Router struct {
	mux     *http.ServeMux
	storage *storage.Storage
	log     *logger.Logger
}

// NewRouter создает роутер и регистрирует эндпоинты
func NewRouter(st *storage.Storage, log *logger.Logger) *Router {
	r := &Router{
		mux:     http.NewServeMux(),
		storage: st,
		log:     log,
	}
	r.register()
	return r
}

// ServeHTTP позволяет использовать Router как http Handler.
func (rt *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rt.mux.ServeHTTP(w, req)
}

// register навешивает хендлеры.
func (rt *Router) register() {
	rt.mux.HandleFunc("POST /tasks", rt.CreateTask)
	rt.mux.HandleFunc("GET /tasks/{id}", rt.GetTaskByID)
	rt.mux.HandleFunc("GET /tasks", rt.ListTasks) // ?status=&limit=&offset=
}
