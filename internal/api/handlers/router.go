package handlers

import (
	"net/http"

	"server/internal/core/logger"
	"server/internal/core/storage"
)

const (
	ListQueryMaxLimit     = 200
	ListQueryDefaultLimit = 50

	ValidationErrorFmt  = "validation error: %s"
	InvalidJSONErrorFmt = "json decoding error: %s"
	StorageErrorFmt     = "storage error: %s"
)

type Router struct {
	mux     http.Handler
	storage *storage.Storage
	log     *logger.Logger
}

// NewRouter создает роутер и регистрирует эндпоинты
func NewRouter(st *storage.Storage, log *logger.Logger) *Router {
	mux := http.NewServeMux()

	r := &Router{
		mux:     log.Middleware(mux),
		storage: st,
		log:     log,
	}

	mux.HandleFunc("POST /tasks", r.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", r.GetTaskByID)
	mux.HandleFunc("GET /tasks", r.ListTasks)

	log.Info("router initialized successfully")
	return r
}

// ServeHTTP позволяет использовать Router как http Handler
func (rt *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rt.mux.ServeHTTP(w, req)
}
