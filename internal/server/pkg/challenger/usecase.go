package challenger

import "github.com/aed86/proof_of_work/internal/server/pkg/challenger/model"

type Usecase interface {
	GetChallenge() (*model.Challenge, error)
	Extract(powHash string) (*model.ChallengeSolveData, error)
	Validate(solvedData model.ChallengeSolveData) error
}
