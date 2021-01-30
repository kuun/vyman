package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/kuun/slog"
)

type _logger struct{}

var log = slog.GetLogger(_logger{})

type debugWriter struct {
	log slog.Logger
}

func (w *debugWriter) Write(p []byte) (n int, err error) {
	p[len(p)-1] = 0
	log.Debug(string(p))
	return len(p), nil
}

type errorWriter struct {
	log slog.Logger
}

func (w *errorWriter) Write(p []byte) (n int, err error) {
	p[len(p)-1] = 0
	log.Error(string(p))
	return len(p), nil
}

func InitLog() {
	gin.DefaultWriter = &debugWriter{log: log}
	gin.DefaultErrorWriter = &errorWriter{log}
}
