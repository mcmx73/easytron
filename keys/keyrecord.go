package keys

import "crypto/ecdsa"

// "crypto/ecdsa"
type WithOption func(*Key)

func NewKey(options ...WithOption) *Key {
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
