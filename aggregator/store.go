package main

import "github.com/yosef-adel/toll-calculator/types"

type MemoryStore struct {
	data map[int]float64
}

func NewMemorySotre() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (s *MemoryStore) Insert(d types.Distance) error {
	s.data[d.OBUID] += d.Value
	return nil
}
