package consumer

import (
	"context"
	"log/slog"
	"math/rand/v2"
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
	c.logger.Info("started consumer")
	// TODO: Implement Consumer functionality.
	c.logger.Info("terminated consumer")
}

// processInvoice constructs invoice data and saves it to storage.
// This function should not be touched.
func (c *Consumer) processInvoice(ID invoice.ID) {
	time.Sleep(1 * time.Second)
	total := currency.USDCents(rand.IntN(1_000_000))
	invoice := invoice.New(ID, total)
	c.storage.Save(ID, invoice)
	c.logger.Debug("processed invoice",
		"ID", ID,
	)
}
