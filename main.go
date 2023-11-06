package main

import "github.com/mcmx73/easytron/rpc"

// TODO crypto module for generate private key and address for Tron/Ethereum
// TODO rpc client module for connect JavaTron/trongid
// TODO rpc client module for connect Ethereum
// TODO rpc server for macOS frontend app
var (
	tronRpcClient *rpc.Client
)

func main() {
	//TODO read config or get from front app
	tronRpcClient = rpc.NewClient(
		rpc.WithUrl("https://api.trongrid.io"),
	)
}
