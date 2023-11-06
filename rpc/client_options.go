package rpc

func WithUrl(url string) ClientWith {
	return func(c *Client) {
		c.nodeUrl = url
	}
}
