package usecase

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/aed86/proof_of_work/internal/pkg/quoter/errors"
	"github.com/aed86/proof_of_work/internal/pkg/quoter/model"
	"github.com/gocarina/gocsv"
)

func (q *quoter) GetQuote() (*model.Quote, error) {
	quotes, err := getAllQuotes()
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	quote := quotes[rand.Intn(len(quotes))]

	return &quote, nil
}

func getAllQuotes() ([]model.Quote, error) {
	absPath, _ := filepath.Abs("./data/quotes.txt")
	file, err := os.OpenFile(absPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.OpenQuoterFileError
	}
	defer file.Close()

	quotes := make([]model.Quote, 0)
	if err := gocsv.UnmarshalFile(file, &quotes); err != nil { // Load clients from file
		return nil, err
	}

	return quotes, nil
}
