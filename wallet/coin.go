package wallet

type CoinDescription struct {
	Id       uint8  `json:"id"`
	Title    string `json:"title"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Coin     bool   `json:"coin"`
	Token    bool   `json:"token"`
}
