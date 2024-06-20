package logger

import (
	"io"
	"log/slog"
	"os"
)

type LogType string

const TextLogType LogType = "TEXT"
const JSONLogType LogType = "JSON"

type LogLevel string

const DebugLogLevel LogLevel = "DEBUG"
const InfoLogLevel LogLevel = "INFO"
const WarnLogLevel LogLevel = "WARN"
const ErrorLogLevel LogLevel = "ERROR"

func Logger(logType LogType, logLevel LogLevel) *slog.Logger {
	level := getLogLevel(logLevel)
	handler := logHandler(logType, level, os.Stdout, nil)
	return slog.New(handler)
}

func getLogLevel(level LogLevel) slog.Level {
	switch level {
	case DebugLogLevel:
		return slog.LevelDebug
	case InfoLogLevel:
		return slog.LevelInfo
	case WarnLogLevel:
		return slog.LevelWarn
	case ErrorLogLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func logHandler(logType LogType, logLevel slog.Level, w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	switch logType {
	case TextLogType:
		return slog.NewTextHandler(w, opts)
	case JSONLogType:
		return slog.NewJSONHandler(w, opts)
	default:
		return slog.NewTextHandler(w, opts)
	}
}
