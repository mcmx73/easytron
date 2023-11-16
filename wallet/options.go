package wallet

import "github.com/mcmx73/easytron/keys"

func WithKeyManager(m *keys.Manager) WithOption {
	return func(w *Manager) {
		w.keyManager = m
	}
}
