package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"server/internal/api/handlers"
	"server/internal/core/logger"
	"server/internal/core/storage"
)

func main() {
	// Инициализация зависимостей
	appLogger := logger.NewLogger("[server] ", 100)
	defer appLogger.Close()

	store := storage.NewStorage()
	r := handlers.NewRouter(store, appLogger)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r, // Router реализует http.Handler
	}

	// Канал для graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		appLogger.Info("Server started on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Ждём сигнал
	<-stop
	appLogger.Info("Shutting down server...")
	if err := srv.Close(); err != nil {
		appLogger.Error("Error during server shutdown: %v", err)
	}
	appLogger.Info("Server stopped")
}
