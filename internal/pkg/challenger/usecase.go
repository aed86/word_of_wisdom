package challenger

import "github.com/aed86/proof_of_work/internal/pkg/challenger/model"

type Usecase interface {
	GetChallenge() (*model.Challenge, error)
	Validate(solvedData model.Solution) error
}
