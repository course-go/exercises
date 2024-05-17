package main

import (
	"log/slog"
	"net/http"

	_ "net/http/pprof"
	"os"

	"github.com/course-go/exercises/06-pprof/internal/handlers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	api := handlers.NewAPI(logger)
	http.HandleFunc("POST /data", api.GenerateData)

	hostname := "localhost:" + port
	logger.Info("started webserver",
		"hostname", hostname,
	)
	http.ListenAndServe(hostname, nil)
}
