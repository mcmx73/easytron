package tronadapter

import "github.com/mcmx73/easytron/rpc"

type WithOption func(*Client)

func NewClient(options ...WithOption) *Client {
	c := &Client{}
	for _, opt := range options {
		opt(c)
	}
	return c
}

type Client struct {
	rpc *rpc.Client
}
