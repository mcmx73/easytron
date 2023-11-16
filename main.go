package main

import (
	"github.com/mcmx73/easytron/frontrpc"
	"github.com/mcmx73/easytron/keys"
	"github.com/mcmx73/easytron/rpc"
	"github.com/mcmx73/easytron/tronadapter"
	"github.com/mcmx73/easytron/wallet"
	"os"
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
	tronAdapter := tronadapter.NewClient(
		tronadapter.WithRpcClient(tronRpcClient),
	)
	keyManager := keys.NewManager()
	walletManager = wallet.NewManager(
		wallet.WithKeyManager(keyManager),
	)
	walletManager.AddCoin(tronAdapter)

	serverOptions := []frontrpc.WithServerOption{
		frontrpc.WithWalletManager(walletManager),
	}

	frontServer := frontrpc.NewServer(serverOptions...)
	err := frontServer.Start()
	if err != nil {
		os.Exit(-1)
	}
	os.Exit(0)
}
