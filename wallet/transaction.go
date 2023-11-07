package wallet

import "time"

type Transaction struct {
	Id            uint64    `json:"id"`
	Hash          string    `json:"hash"`
	Incoming      bool      `json:"incoming"`
	Outgoing      bool      `json:"outgoing"`
	CreatedAt     time.Time `json:"created_at"`
	ConfirmedAt   time.Time `json:"confirmed_at"`
	Confirmations uint64    `json:"confirmations"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Amount        Amount    `json:"amount"`
	Fee           Amount    `json:"fee"`
}
