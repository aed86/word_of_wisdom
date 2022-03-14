package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	host   string
	client http.Client
}

func NewClient() *Client {
	return &Client{
		host: "0.0.0.0:8000", // TODO: move to env
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
