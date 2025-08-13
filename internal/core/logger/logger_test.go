package logger

import (
	"testing"
	"time"
)

func TestLogger_InfoAndError(t *testing.T) {
	logger := NewLogger("[test] ", 10)
	defer logger.Close()

	logger.Info("info message")
	logger.Error("error message")

	time.Sleep(50 * time.Millisecond)
}
