package usecase

import (
	"math/big"

	"github.com/aed86/proof_of_work/internal"
	challenger2 "github.com/aed86/proof_of_work/internal/pkg/challenger"
)

type challenger struct {
	logger            internal.Logger
	leadingZerosCount int
	target            *big.Int
}

func NewChallenger(logger internal.Logger) *challenger {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-challenger2.LeadingZerosCount*8))

	return &challenger{
		logger:            logger,
		target:            target,
		leadingZerosCount: challenger2.LeadingZerosCount,
	}
}
