package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"math/big"
	"time"

	"github.com/aed86/proof_of_work/internal/server/pkg/challenger/errors"
	"github.com/aed86/proof_of_work/internal/server/pkg/challenger/model"
	"github.com/aed86/proof_of_work/internal/utils"
)

//Validate Challenge Validator
func (c *challenger) Validate(solvedData model.ChallengeSolveData) error {
	// check challenge ttl
	if time.Now().Sub(solvedData.Time) > c.ttl {
		return errors.OutdatedChallenge
	}

	// compare expected and provided pow hashes
	expectedHash := c.getHash(solvedData.Data, solvedData.Time)
	if !hmac.Equal(expectedHash, solvedData.AuthCode) {
		return errors.InvalidChallenge
	}

	// check expected prepend zeros in hash // TODO: check
	challengeHash := sha256.Sum256(append(solvedData.Data, utils.Int64ToBytes(solvedData.Nonce)...))

	//TODO: find another way to check zeros
	var challengeHashInt big.Int
	if challengeHashInt.SetBytes(challengeHash[:]).Cmp(c.target) != -1 {
		return errors.InvalidChallenge
	}

	return nil
}
