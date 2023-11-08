package wallet

type CoinId string
type CoinDescription struct {
	Id       CoinId `json:"id"`
	Title    string `json:"title"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Coin     bool   `json:"coin"`
	Token    bool   `json:"token"`
}
