package logger

import (
	"log/slog"
	"os"
)

func NewDebug() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
