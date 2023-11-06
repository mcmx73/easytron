package rpc

import "net/http"

type ClientWith func(*Client)

func NewClient(options ...ClientWith) (c *Client) {
	c = &Client{}
	for _, opt := range options {
		opt(c)
	}
	return c
}

type Client struct {
	nodeUrl    string
	httpClient *http.Client
}

func (c *Client) Init() error {
	return nil
}
