package rpc

func WithUrl(url string) WithOption {
	return func(c *Client) {
		c.nodeUrl = url
	}
}
