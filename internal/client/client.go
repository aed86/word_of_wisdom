package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aed86/proof_of_work/internal"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
	model2 "github.com/aed86/proof_of_work/internal/pkg/quoter/model"
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

func (c *Client) GetChallenge() (*model.Challenge, error) {
	req, err := http.NewRequest("GET", c.getChallengeUri(), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.do(req)
	if err != nil {
		return nil, err
	}

	challenge := &model.Challenge{}
	err = json.Unmarshal(body, challenge)
	if err != nil {
		return nil, err

	}

	return challenge, nil
}

func (c *Client) GetQuote(powHeader string) (*model2.Quote, error) {
	req, err := http.NewRequest("GET", c.getQuoteUri(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(internal.PowHeaderName, powHeader)
	body, err := c.do(req)
	if err != nil {
		return nil, err
	}

	quote := &model2.Quote{}
	err = json.Unmarshal(body, quote)
	if err != nil {
		return nil, err
	}

	return quote, nil
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

func (c *Client) getChallengeUri() string {
	return fmt.Sprintf("http://%s%s", c.host, internal.GetChallengeEndpoint)
}

func (c *Client) getQuoteUri() string {
	return fmt.Sprintf("http://%s%s", c.host, internal.GetQuoteEndpoint)
}
