package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/course-go/exercises/09-databases/internal/todos"
)

const (
	portEnv = "PORT"
)

const defaultPort = "8080"

func main() {
	repository := todos.NewRepository()
	mux := todos.NewRouter(repository)
	port := os.Getenv(portEnv)
	if port == "" {
		slog.Info("defaulting variable value",
			"variable", portEnv,
			"value", defaultPort,
		)
		port = defaultPort
	}

	hostname := fmt.Sprintf(":%s", port)
	slog.Info("starting server",
		"hostname", hostname,
	)
	err := http.ListenAndServe(hostname, mux)
	slog.Error("failed running server",
		"error", err,
	)
}
