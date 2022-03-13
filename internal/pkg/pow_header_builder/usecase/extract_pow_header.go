package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
)

func (*powHeaderBuilder) Extract(powHeader string) (*model.Solution, error) {
	parts := strings.Split(powHeader, ":")
	if len(parts) < 4 {
		return nil, fmt.Errorf("wrong header format")
	}

	nonce, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return nil, err
	}

	timestamp, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		return nil, err
	}

	return &model.Solution{
		Challenge: []byte(parts[0]),
		Nonce:     nonce,
		Timestamp: timestamp,
	}, nil
}
