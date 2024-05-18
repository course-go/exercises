package invoice

import "github.com/course-go/04-concurrency/internal/currency"

type ID uint64

type Invoice struct {
	ID    ID
	Total currency.USDCents
}

func New(ID ID, total currency.USDCents) Invoice {
	return Invoice{
		ID:    ID,
		Total: total,
	}
}
