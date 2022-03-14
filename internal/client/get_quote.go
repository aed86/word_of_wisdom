package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aed86/proof_of_work/internal"
	model2 "github.com/aed86/proof_of_work/internal/pkg/quoter/model"
)

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

func (c *Client) getQuoteUri() string {
	return fmt.Sprintf("http://%s%s", c.host, internal.GetQuoteEndpoint)
}
