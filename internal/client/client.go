package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	address string
	client  http.Client
}

func NewClient(address string) *Client {
	return &Client{
		address: address,
		client: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("client error: %s", string(body))
	}

	return body, nil
}
