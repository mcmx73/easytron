package tronadapter

import (
	"github.com/mcmx73/easytron/keys"
	"github.com/mcmx73/easytron/wallet"
)

func (c *Client) GetCoins() (coins []*wallet.CoinDescription) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) CreateNewAddress(privateKey *keys.Key) (address string, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetAddressBalance(address string) (amounts map[string]wallet.Amount, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetAddressTransactions(address string) (transactions []*wallet.Transaction, err error) {
	//TODO implement me
	panic("implement me")
}
