package invoice

import (
	"log/slog"
)

// Generator is used for generating new invoice IDs.
type Generator struct {
	logger        *slog.Logger
	lastInvoiceID ID
}

func NewGenerator(logger *slog.Logger) *Generator {
	return &Generator{
		lastInvoiceID: 0,
		logger:        logger,
	}
}

func (g *Generator) ID() ID {
	g.lastInvoiceID++
	return g.lastInvoiceID
}
