package utils

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"server/internal/core/logger"
)

// RunHTTPServer запускает сервер и делает graceful shutdown по SIGINT/SIGTERM.
// timeout - сколько ждать завершения активных запросов.
func RunHTTPServer(srv *http.Server, log *logger.Logger, timeout time.Duration) {
	// Стартуем слушателя в горутине
	go func() {
		log.Info("server starting on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server error: %v", err)
		}
	}()

	// Ждём сигнал завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop
	log.Info("shutdown signal received: %v", sig)

	// Контекст с таймаутом для плавной остановки
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Останавливаем сервер: перестаёт принимать новые соединения,
	// ждёт активные запросы до timeout
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("graceful shutdown error: %v", err)
	} else {
		log.Info("server stopped gracefully")
	}
}
