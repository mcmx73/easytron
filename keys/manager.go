package keys

import "sync"

type WithManagerOption func(*Manager)

func NewManager(options ...WithManagerOption) *Manager {
	m := &Manager{
		keys: make(map[string]*Key),
	}
	for _, opt := range options {
		opt(m)
	}
	return m
}

type Manager struct {
	keyMux sync.RWMutex
	keys   map[string]*Key
}
