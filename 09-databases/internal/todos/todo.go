package todos

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID  `json:"id,omitzero"`
	Description string     `json:"description,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitzero"`
	CreatedAt   time.Time  `json:"created_at,omitzero"`
	UpdatedAt   *time.Time `json:"updated_at,omitzero"`
}
