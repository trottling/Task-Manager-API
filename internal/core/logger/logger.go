package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogEvent описывает одно сообщение лога
type LogEvent struct {
	Level   string
	Message string
	Time    time.Time
}

// Logger — асинхронный логгер через канал
type Logger struct {
	ch     chan LogEvent
	quit   chan struct{}
	prefix string
}

// NewLogger создаёт новый логгер с буфером буфера bufferSize.
func NewLogger(prefix string, bufferSize int) *Logger {
	l := &Logger{
		ch:     make(chan LogEvent, bufferSize),
		quit:   make(chan struct{}),
		prefix: prefix,
	}

	go l.Worker()
	return l
}

// Worker — горутина, которая слушает канал и пишет логи
func (l *Logger) Worker() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	for {
		select {
		case e := <-l.ch:
			ts := e.Time.Format("2006-01-02 15:04:05")
			logger.Printf("[%s] %s%s: %s\n", e.Level, l.prefix, ts, e.Message)
		case <-l.quit:
			return
		}
	}
}

// Info пишет информационное сообщение
func (l *Logger) Info(msg string, args ...any) {
	l.push("INFO", msg, args...)
}

// Error пишет сообщение об ошибке
func (l *Logger) Error(msg string, args ...any) {
	l.push("ERROR", msg, args...)
}

func (l *Logger) push(level, msg string, args ...any) {
	select {
	case l.ch <- LogEvent{
		Level:   level,
		Message: fmt.Sprintf(msg, args...),
		Time:    time.Now(),
	}:
	default:
		// Канал переполнен — можно дропать или блокировать
	}
}

// Close завершает работу логгера
func (l *Logger) Close() {
	close(l.quit)
}
