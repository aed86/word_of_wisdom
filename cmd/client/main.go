package main

import (
	"encoding/base64"
	"log"

	"github.com/aed86/proof_of_work/internal/client"
	challenger_usecase "github.com/aed86/proof_of_work/internal/pkg/challenger/usecase"
	"github.com/aed86/proof_of_work/internal/pkg/pow_header_builder/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	var (
		c      = client.NewClient()
		e      = echo.New()
		logger = e.Logger
	)

	logger.Print("Started")
	challenge, err := c.GetChallenge()
	if err != nil {
		logger.Fatal(err)
	}

	data := base64.StdEncoding.EncodeToString(challenge.ChallengeData)
	logger.Print("Challenge received:", data)

	powHeaderBuilder := usecase.NewPowHeaderBuilder()
	challenger := challenger_usecase.NewChallenger(logger)

	solution := challenger.Solve(*challenge)
	logger.Print("Solution: ", base64.StdEncoding.EncodeToString(solution.Hash))

	powHeader := powHeaderBuilder.Build(*solution)
	logger.Print("pow-header prepared: ", powHeader)

	quote, err := c.GetQuote(powHeader)
	if err != nil {
		log.Fatal(err)
	}

	logger.Print(quote)
	logger.Print("Finished")
}
