package usecase

import (
	"crypto/sha256"
	"math"
	"math/big"

	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/utils"
)

func (c *challenger) Solve(challenge model.Challenge) *model.Solution {
	target := c.target
	solution := make([]byte, 0, len(challenge.Challenge))
	var (
		hashToCompare big.Int
		nonce         int64
		hash          [32]byte
	)

	for nonce < math.MaxInt64 {
		solution = append(solution, challenge.Challenge...)
		solution = append(solution, utils.Int64ToBytes(nonce)...)

		hash = sha256.Sum256(solution)
		hashToCompare.SetBytes(hash[:])

		if hashToCompare.Cmp(target) == -1 {
			break
		}

		solution = solution[:0]
		nonce++
	}

	return &model.Solution{
		Nonce:    nonce,
		Solution: challenge.Challenge,
	}
}
