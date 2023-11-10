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
	mux          sync.RWMutex
	clients      map[CoinId]Blockchain
	coinsById    map[CoinId]*CoinDescription
	coinsByTitle map[string]*CoinDescription
}

func (m *Manager) AddCoin(client Blockchain) {
	m.mux.Lock()
	defer m.mux.Unlock()
	coins := client.GetCoins()
	for _, coin := range coins {
		m.coinsById[coin.Id] = coin
		m.coinsByTitle[coin.Title] = coin
		m.clients[coin.Id] = client
	}
}
