package http

import (
	"net/http"
)

type Client struct {
	client *http.Client
	ua     string
}

func New(client *http.Client) *Client {
	return &Client{
		client: client,
		ua:     "gopwd-http-client",
	}
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.client = client
}
