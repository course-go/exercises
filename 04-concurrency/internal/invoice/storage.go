package invoice

import (
	"context"
	"log/slog"
	"time"
)

// Storage represents an in-memory storage for storing
// invoice IDs with their respective invoices.
type Storage struct {
	logger   *slog.Logger
	invoices map[ID]Invoice
}

func NewStorage(logger *slog.Logger) *Storage {
	return &Storage{
		logger: logger,
	}
}

func (s *Storage) Save(ID ID, invoice Invoice) {
	s.invoices[ID] = invoice
}

// Observer observes storage state and logs the amount of stored items.
func (s *Storage) RunObserver(ctx context.Context) {
	t := time.NewTicker(3 * time.Second)
	count := len(s.invoices)
	for {
		select {
		case <-t.C:
			updatedCount := len(s.invoices)
			s.logger.Debug("stored items",
				"count", updatedCount,
				"delta", updatedCount-count,
			)
			count = updatedCount
		case <-ctx.Done():
			s.logger.Info("terminated observer")
			return
		}
	}
}
