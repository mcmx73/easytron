package wallet

import "sync"

type WithOption func(*Manager)

func NewManager(options ...WithOption) *Manager {
	m := &Manager{
		clients: make(map[CoinId]Blockchain),
	}
	for _, opt := range options {
		opt(m)
	}
	return m
}

type Manager struct {
	mux     sync.RWMutex
	clients map[CoinId]Blockchain
}

func (m *Manager) AddCoin(client Blockchain) {
	m.mux.Lock()
	defer m.mux.Unlock()
	coins := client.GetCoins()
	for _, coin := range coins {
		m.clients[coin.Id] = client
	}
}
