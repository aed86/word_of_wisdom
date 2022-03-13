package usecase

import (
	"math/big"
	"time"
)

type challenger struct {
	zeroBytesCount int
	target         *big.Int
	secretKey      []byte
	ttl            time.Duration
}

func NewChallenger() *challenger {
	return &challenger{}
}
