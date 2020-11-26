package main

import (
	"sync"
)

type store struct {
	data  map[uint64]uint64
	mutex sync.Mutex
}

func newStore() *store {
	return &store{
		data: make(map[uint64]uint64),
	}
}

func (s *store) Get(n uint64) uint64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.data[n]
}

func (s *store) Has(n uint64) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.data[n]
	return ok
}

func (s *store) Set(n, value uint64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[n] = value
}
