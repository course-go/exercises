package producer

import (
	"context"
	"log/slog"
	"time"

	"github.com/course-go/04-concurrency/internal/invoice"
)

type Producer struct {
	logger    *slog.Logger
	generator *invoice.Generator
}

func New(logger *slog.Logger, generator *invoice.Generator) Producer {
	return Producer{
		logger:    logger,
		generator: generator,
	}
}

func (p *Producer) Start(ctx context.Context, _ chan<- invoice.ID) {
	slog.Info("started producer")
	// TODO: Implement Producer functionality.
	slog.Info("terminated producer")
}

// generateInvoiceID generates new invoice ID using the invoice generator.
// This function should not be touched.
func (p *Producer) generateInvoiceID() invoice.ID {
	time.Sleep(250 * time.Millisecond)
	ID := p.generator.ID()
	slog.Debug("generated invoice",
		"ID", ID,
	)
	return ID
}
