package usecase

import (
	"math/big"
	"time"

	challenger2 "github.com/aed86/proof_of_work/internal/pkg/challenger"
)

type challenger struct {
	leadingZerosCount int
	target            *big.Int
	ttl               time.Duration
}

func NewChallenger() *challenger {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-challenger2.LeadingZerosCount*8))

	return &challenger{
		target:            target,
		leadingZerosCount: challenger2.LeadingZerosCount,
		ttl:               time.Second * 10,
	}
}
