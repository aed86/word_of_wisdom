package usecase

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
)

func (*powHeaderBuilder) Extract(powHeader string) (*model.Solution, error) {
	parts := strings.Split(powHeader, ":")
	if len(parts) < 2 {
		return nil, fmt.Errorf("wrong header format")
	}

	challenge, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}

	nonce, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return &model.Solution{
		Solution: challenge,
		Nonce:    nonce,
	}, nil
}
