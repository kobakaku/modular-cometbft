package store

import "sync/atomic"

type Store struct {
	height atomic.Uint64
}

func New() *Store {
	return &Store{}
}

func (s *Store) SetHeight(height uint64) {
	s.height.Store(height)
}

func (s *Store) Height() uint64 {
	return s.height.Load()
}
