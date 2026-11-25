package main

import (
	"fmt"
	"sync"
)

type Store struct {
	data map[string]int
	mu   sync.Mutex
}

func NewStore() *Store {
	return &Store{data: make(map[string]int)}
}

func main() {
	mapper := make(map[string]int)
	store := NewStore()
	mapper["KTA"] = 27
	store.Put("KTA", 27)
	fmt.Println("Hello, World!")
}

func (s *Store) Get(key string) (int, bool) {
	s.mu.Lock()
	_, ok := s.data[key]
	if ok {
		s.mu.Unlock()
		return s.data[key], true
	}
	s.mu.Unlock()
	return 0, false
}

func (s *Store) Put(key string, val int) {
	s.mu.Lock()
	s.data[key] = val
	s.mu.Unlock()
}

func (s *Store) Del(key string) bool {
	s.mu.Lock()
	_, found := s.data[key]
	if found {
		delete(s.data, key)
		s.mu.Unlock()
		return true
	}
	s.mu.Unlock()
	return false

}
