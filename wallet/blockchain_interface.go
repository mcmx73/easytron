package wallet

type Blockchain interface {
	GetCoins() (coins []*CoinDescription)
	CreateNewAddress() (privateKey, publicKey, address string, err error)
	GetAddressBalance(address string) (amounts map[string]Amount, err error)
	GetAddressTransactions(address string) (transactions []*Transaction, err error)
}
