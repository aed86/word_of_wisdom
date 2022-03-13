package usecase

import (
	"crypto/rand"
	"time"

	challenger2 "github.com/aed86/proof_of_work/internal/server/pkg/challenger"
	"github.com/aed86/proof_of_work/internal/server/pkg/challenger/model"
)

func (c *challenger) GetChallenge() (*model.Challenge, error) {
	data := make([]byte, 16)
	_, err := rand.Read(data)
	if err != nil {
		return nil, err
	}

	hash := c.getHash(data, time.Now())

	return &model.Challenge{
		Data:             data,
		AuthCodeHash:     hash,
		PrependZeroCount: challenger2.PrependZeroCount,
		Time:             time.Now(),
	}, nil
}
