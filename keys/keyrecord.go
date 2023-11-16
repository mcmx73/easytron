package keys

import "crypto/ecdsa"

// "crypto/ecdsa"
type WithKeyOption func(*Key)

func NewKey(options ...WithKeyOption) *Key {
	k := &Key{}
	for _, opt := range options {
		opt(k)
	}
	return k
}

type Key struct {
	PublicKey  string
	PrivateKey string
	privateKey *ecdsa.PrivateKey
}
