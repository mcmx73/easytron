package wallet

import "github.com/mcmx73/easytron/keys"

type Blockchain interface {
	GetCoins() (coins []*CoinDescription)
	CreateNewAddress(privateKey *keys.Key) (address string, err error)
	GetAddressBalance(address string) (amounts map[string]Amount, err error)
	GetAddressTransactions(address string) (transactions []*Transaction, err error)
}
