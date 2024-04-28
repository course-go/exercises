package main

import (
	"fmt"
	"os"

	"github.com/course-go/exercises/observability/internal/todos"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	databaseURLEnv = "DATABASE_URL"
	portEnv        = "PORT"
)

const defaultPort = "8080"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		return
	}

	databaseURL := os.Getenv(databaseURLEnv)
	if databaseURL == "" {
		logger.Fatal("environment variable not set",
			zap.String("variable", databaseURLEnv),
		)
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		logger.Fatal("could not open connection to database",
			zap.Error(err),
		)
	}

	repository := todos.NewRepository(logger, db)
	router := todos.NewRouter(logger, repository)
	port := os.Getenv(portEnv)
	if port == "" {
		logger.Info("defaulting variable value",
			zap.String("variable", portEnv),
			zap.String("value", defaultPort),
		)
		port = defaultPort
	}

	hostname := fmt.Sprintf(":%s", port)
	logger.Info("starting server",
		zap.String("hostname", hostname),
	)

	err = router.Run(hostname)
	if err != nil {
		logger.Fatal("failed starting server",
			zap.Error(err),
		)
	}
}
