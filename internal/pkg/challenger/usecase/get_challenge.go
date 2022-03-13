package usecase

import (
	"crypto/rand"
	"time"

	challenger2 "github.com/aed86/proof_of_work/internal/pkg/challenger"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
)

func (c *challenger) GetChallenge() (*model.Challenge, error) {
	challengeData := make([]byte, 16)
	_, err := rand.Read(challengeData)
	if err != nil {
		return nil, err
	}

	return &model.Challenge{
		ChallengeData: challengeData,
		Timestamp:     time.Now().UnixNano(),
		LeadingZeros:  challenger2.LeadingZerosCount,
	}, nil
}
