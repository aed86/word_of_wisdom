package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aed86/proof_of_work/internal"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
)

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

func (c *Client) getChallengeUri() string {
	return fmt.Sprintf("http://%s%s", c.address, internal.GetChallengeEndpoint)
}
