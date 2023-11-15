package wallet

import (
	"github.com/mcmx73/easytron/keys"
	"sync"
)

type KeyStorage struct {
	mux  sync.RWMutex
	keys map[uint64]*keys.Key
}

func (s *KeyStorage) init() {
	s.keys = make(map[uint64]*keys.Key)
}

func (s *KeyStorage) AddKey(keyId uint64, key *keys.Key) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.keys[keyId] = key
}

func (s *KeyStorage) GetKey(keyId uint64) (key *keys.Key, found bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	key, found = s.keys[keyId]
	return key, found
}
