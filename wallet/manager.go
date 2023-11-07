package wallet

import "sync"

type WithOption func(*Manager)

func NewManager(options ...WithOption) *Manager {
	m := &Manager{
		clients: make(map[string]Blockchain),
	}
	for _, opt := range options {
		opt(m)
	}
	return m
}

type Manager struct {
	mux     sync.RWMutex
	clients map[string]Blockchain
}
