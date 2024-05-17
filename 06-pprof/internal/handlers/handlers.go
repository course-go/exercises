package handlers

import (
	"log/slog"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/course-go/exercises/06-pprof/internal/business"
	"github.com/course-go/exercises/06-pprof/internal/storage"
)

const defaultCount = "128"

type API struct {
	logger   *slog.Logger
	business *business.Business
}

func NewAPI(logger *slog.Logger) *API {
	return &API{
		logger:   logger,
		business: business.New(logger),
	}
}

func (a *API) GenerateData(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	countValue := r.URL.Query().Get("count")
	if countValue == "" {
		countValue = defaultCount
	}

	count, err := strconv.ParseInt(countValue, 10, 64)
	if err != nil || count <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := storage.ID(rand.Int())
	a.logger.Debug("received generate data request",
		"ID", id,
	)
	time.Sleep(150 * time.Millisecond)
	a.business.GenerateData(id, count)
	w.WriteHeader(http.StatusOK)
}
