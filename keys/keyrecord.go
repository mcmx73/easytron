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
	address    string
	pubKey     string
	pkString   string
	privateKey *ecdsa.PrivateKey
}
