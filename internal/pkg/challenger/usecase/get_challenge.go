package usecase

import (
	"crypto/rand"

	challenger2 "github.com/aed86/proof_of_work/internal/pkg/challenger"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
)

func (c *challenger) GetChallenge() (*model.Challenge, error) {
	challengeData := make([]byte, 32)
	_, err := rand.Read(challengeData)
	if err != nil {
		return nil, err
	}

	return &model.Challenge{
		Challenge:    challengeData,
		LeadingZeros: challenger2.LeadingZerosCount,
	}, nil
}
