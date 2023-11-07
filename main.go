package main

import (
	"github.com/mcmx73/easytron/rpc"
	"github.com/mcmx73/easytron/tronclient"
	"github.com/mcmx73/easytron/wallet"
)

// TODO crypto module for generate private key and address for Tron/Ethereum
// TODO rpc client module for connect JavaTron/trongid
// TODO rpc client module for connect Ethereum
// TODO rpc server for macOS frontend app
var (
	walletManager *wallet.Manager
)

func main() {
	//TODO read config or get from front app
	tronRpcClient := rpc.NewClient(
		rpc.WithUrl("https://api.trongrid.io"),
	)
	tronClient := tronclient.NewClient(
		tronclient.WithRpcClient(tronRpcClient),
	)
}
