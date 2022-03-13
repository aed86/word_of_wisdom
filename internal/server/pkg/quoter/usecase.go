package quoter

import "github.com/aed86/proof_of_work/internal/server/pkg/quoter/model"

type Usecase interface {
	GetQuote() (*model.Quote, error)
}
