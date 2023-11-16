package tronadapter

import "github.com/mcmx73/easytron/rpc"

func WithRpcClient(rpc *rpc.Client) WithOption {
	return func(c *Client) {
		c.rpc = rpc
	}
}
