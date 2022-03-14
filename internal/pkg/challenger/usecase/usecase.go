package usecase

import (
	"math/big"

	challenger2 "github.com/aed86/proof_of_work/internal/pkg/challenger"
	"github.com/labstack/echo/v4"
)

type challenger struct {
	logger            echo.Logger
	leadingZerosCount int
	target            *big.Int
}

func NewChallenger(logger echo.Logger) *challenger {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-challenger2.LeadingZerosCount*8))

	return &challenger{
		logger:            logger,
		target:            target,
		leadingZerosCount: challenger2.LeadingZerosCount,
	}
}
