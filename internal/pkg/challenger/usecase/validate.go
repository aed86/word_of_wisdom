package usecase

import (
	"crypto/sha256"
	"encoding/base64"
	"math/big"

	"github.com/aed86/proof_of_work/internal/pkg/challenger/errors"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/utils"
)

//Validate Solution Validator
func (c *challenger) Validate(solve model.Solution) error {
	c.logger.Print("Target: ", c.target)

	// get hash from summed challenge and nonce
	var solutionHashToCompare big.Int
	solution := make([]byte, 0, len(solve.Solution))
	solution = append(solution, solve.Solution...)
	solution = append(solution, utils.Int64ToBytes(solve.Nonce)...)

	solutionHash := sha256.Sum256(solution)
	solutionHashToCompare.SetBytes(solutionHash[:])

	c.logger.Print("Solution: ", base64.StdEncoding.EncodeToString(solutionHash[:]))

	if solutionHashToCompare.Cmp(c.target) != -1 {
		return errors.InvalidSolve
	}

	return nil
}
