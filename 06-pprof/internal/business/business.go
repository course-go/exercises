package business

import (
	"log/slog"
	"math/rand"
	"time"

	"github.com/course-go/exercises/06-pprof/internal/storage"
)

const dataSize = 1024

type Business struct {
	logger  *slog.Logger
	storage *storage.InMemory
}

func New(logger *slog.Logger) *Business {
	return &Business{
		logger:  logger,
		storage: storage.NewInMemory(logger),
	}
}

func (b *Business) GenerateData(id storage.ID, count int64) {
	data := make([][]int, 0)
	for range count {
		numbers := make([]int, 0)
		for range dataSize {
			number := rand.Int()
			numbers = append(numbers, number)
		}

		data = append(data, numbers)
	}

	time.Sleep(250 * time.Millisecond)
	b.logger.Debug("generated data",
		"ID", id,
	)
	b.storage.Store(id, data)
}
