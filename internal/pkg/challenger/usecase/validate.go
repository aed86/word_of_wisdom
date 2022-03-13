package usecase

import (
	"crypto/sha256"
	"math/big"
	"time"

	"github.com/aed86/proof_of_work/internal/pkg/challenger/errors"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/utils"
)

//Validate Challenge Validator
func (c *challenger) Validate(solve model.Solution) error {
	// check challenge ttl
	if time.Now().Sub(time.Unix(solve.Timestamp, 0)) > c.ttl {
		return errors.OutdatedChallenge
	}

	// get hash from summed challenge and nonce
	challengeHash := sha256.Sum256(append(solve.Challenge, utils.Int64ToBytes(solve.Nonce)...))
	var challengeHashInt big.Int
	if challengeHashInt.SetBytes(challengeHash[:]).Cmp(c.target) != -1 {
		return errors.InvalidSolve
	}

	return nil
}
