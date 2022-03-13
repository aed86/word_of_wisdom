package usecase

import (
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/utils"
)

func (c *challenger) Solve(challenge model.Challenge) *model.Solution {
	target := c.target

	solution := make([]byte, 0, len(challenge.ChallengeData))
	var (
		hashToCompare big.Int
		nonce         int64
	)
	for nonce < math.MaxInt64 {
		solution = append(solution, challenge.ChallengeData...)
		solution = append(solution, utils.Int64ToBytes(nonce)...)

		hash := sha256.Sum256(solution)
		hashToCompare.SetBytes(hash[:])

		if hashToCompare.Cmp(target) == -1 {
			break
		}

		solution = solution[:0]
		nonce++
	}
	fmt.Println("Prepared solution: %s", string(solution))

	return &model.Solution{
		Nonce:        nonce,
		Challenge:    challenge.ChallengeData,
		Timestamp:    time.Now().UnixNano(),
		LeadingZeros: challenge.LeadingZeros,
	}
}
