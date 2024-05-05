package todo

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt *time.Time
}
