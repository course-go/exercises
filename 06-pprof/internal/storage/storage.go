package storage

import (
	"log/slog"
	"sync"
	"time"
)

type ID int
type Data [][]int

type InMemory struct {
	logger *slog.Logger
	mu     sync.Mutex
	data   map[ID]Data
}

func NewInMemory(logger *slog.Logger) *InMemory {
	return &InMemory{
		logger: logger,
		data:   make(map[ID]Data),
	}
}

func (m *InMemory) Store(id ID, data Data) {
	m.mu.Lock()
	defer m.mu.Unlock()
	time.Sleep(100 * time.Millisecond)
	m.data[id] = data
	m.logger.Debug("saved data",
		"ID", id,
	)
	go func() {
		time.AfterFunc(250*time.Millisecond, func() {
			m.Invalidate(id)
		})
	}()
}

func (m *InMemory) Invalidate(id ID) {
	m.mu.Lock()
	defer m.mu.Unlock()
	time.Sleep(50 * time.Millisecond)
	delete(m.data, id)
	m.logger.Debug("invalidated data",
		"ID", id,
	)
}
