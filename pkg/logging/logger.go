package logging

import (
	"context"
	"github.com/csyezheng/memcard/pkg/utils"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// contextKey is a private string type to prevent collisions in the context map.
type contextKey string

// loggerKey points to the value in the context where the logger is stored.
const loggerKey = contextKey("logger")

var (
	defaultLogger     *slog.Logger
	defaultLoggerOnce sync.Once
)

const (
	levelDebug   = "DEBUG"
	levelInfo    = "INFO"
	levelWarning = "WARNING"
	levelError   = "ERROR"
)

func levelToSlogLevel(s string) slog.Leveler {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case levelDebug:
		return slog.LevelDebug
	case levelInfo:
		return slog.LevelInfo
	case levelWarning:
		return slog.LevelWarn
	case levelError:
		return slog.LevelError
	}
	return slog.LevelWarn
}

func NewLogger(level string, form string) *slog.Logger {
	logOutput := os.Stdout
	output := utils.GetEnv("LOG_OUTPUT", "Stdout")
	if output == "Stderr" {
		logOutput = os.Stderr
	}
	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove time.
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	opts := slog.HandlerOptions{AddSource: true, Level: levelToSlogLevel(level), ReplaceAttr: replace}
	var logHandler slog.Handler
	switch form {
	case "json":
		logHandler = slog.NewJSONHandler(logOutput, &opts)
	case "text":
		logHandler = slog.NewTextHandler(logOutput, &opts)
	}
	return slog.New(logHandler)
}

func NewLoggerFromEnv() *slog.Logger {
	level := utils.GetEnv("LOG_LEVEL", "INFO")
	form := utils.GetEnv("LOG_FORM", "text")
	return NewLogger(level, form)
}

func DefaultLogger() *slog.Logger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewLoggerFromEnv()
	})
	return defaultLogger
}

// AttachLogger creates a new context with the provided logger attached.
func AttachLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext returns the logger stored in the context. If no such logger
// exists, a default logger is returned.
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}
	return DefaultLogger()
}
