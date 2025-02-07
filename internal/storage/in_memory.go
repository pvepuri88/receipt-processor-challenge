package storage

import (
	"errors"
	"sync"
)

type Store interface {
	Save(id string, points int) error
	Get(id string) (int, error)
}

type InMemoryStore struct {
	mu     sync.RWMutex
	values map[string]int
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{values: make(map[string]int)}
}

func (s *InMemoryStore) Save(id string, points int) error {
	s.mu.Lock()
	s.values[id] = points
	s.mu.Unlock()
	return nil
}

func (s *InMemoryStore) Get(id string) (int, error) {
	s.mu.RLock()
	val, ok := s.values[id]
	s.mu.RUnlock()
	if !ok {
		return 0, errors.New("not found")
	}
	return val, nil
}

