package katas

import (
	"log/slog"

	logging "github.com/course-go/exercises/02-katas/internal/katas/slog"
)

// Shadowing create a new logger depending on environment.
// However, if you run it it panics.
// Find out why and fix it.
func Shadowing(environment string) {
	// Slog is standard library structured logger.
	// It will be explored in following lectures.
	var logger *slog.Logger
	if environment == "testing" {
		logger := logging.TestingLogger()
		logger.Debug("initialized testing logger")
	} else {
		logger := logging.ProductionLogger()
		logger.Debug("initialized production logger")
	}

	logger.Debug("initialized all dependencies")
}
