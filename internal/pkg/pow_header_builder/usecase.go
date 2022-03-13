package pow_header_builder

import "github.com/aed86/proof_of_work/internal/pkg/challenger/model"

type Usecase interface {
	Build(solution model.Solution) string
	Extract(powHeader string) (*model.Solution, error)
}
