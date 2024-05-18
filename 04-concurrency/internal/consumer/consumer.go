package consumer

import (
	"context"
	"log/slog"
	"math/rand"
	"time"

	"github.com/course-go/04-concurrency/internal/currency"
	"github.com/course-go/04-concurrency/internal/invoice"
)

type Consumer struct {
	logger  *slog.Logger
	storage *invoice.Storage
}

func New(logger *slog.Logger, storage *invoice.Storage) Consumer {
	return Consumer{
		logger:  logger,
		storage: storage,
	}
}

func (c *Consumer) Start(_ context.Context, _ <-chan invoice.ID) {
	slog.Info("started consumer")
	// TODO: Implement Consumer functionality.
	slog.Info("terminated consumer")
}

// processInvoice constructs invoice data and saves it to storage.
// This function should not be touched.
func (c *Consumer) processInvoice(ID invoice.ID) {
	time.Sleep(1 * time.Second)
	total := currency.USDCents(rand.Int() % 1_000_000)
	invoice := invoice.New(ID, total)
	c.storage.Save(ID, invoice)
	slog.Debug("processed invoice",
		"ID", ID,
	)
}
