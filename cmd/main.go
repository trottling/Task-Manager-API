package main

import (
	"net/http"
	"server/internal/utils"
	"time"

	"server/internal/api/handlers"
	"server/internal/core/logger"
	"server/internal/core/storage"
)

func main() {
	// Инициализируем логгер
	log := logger.NewLogger("[server] ", 100)
	defer log.Close()

	// Инициализируем хранилище
	st := storage.NewStorage()

	// Инициализируем роутер с хендлерами
	r := handlers.NewRouter(st, log)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	utils.RunHTTPServer(srv, log, 15*time.Second)
}
