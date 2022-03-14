package main

import (
	"github.com/aed86/proof_of_work/internal"
	"github.com/aed86/proof_of_work/internal/pkg/challenger/usecase"
	pow_header_builder_usecase "github.com/aed86/proof_of_work/internal/pkg/pow_header_builder/usecase"
	quoter_usecase "github.com/aed86/proof_of_work/internal/pkg/quoter/usecase"
	"github.com/aed86/proof_of_work/internal/server/handlers/get_challenge"
	"github.com/aed86/proof_of_work/internal/server/handlers/get_quote"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	challengerUsecase := usecase.NewChallenger(e.Logger)
	quoterUsecase := quoter_usecase.NewQuoter()
	powHeaderBuilderUsecase := pow_header_builder_usecase.NewPowHeaderBuilder()

	getChallengerHandler := get_challenge.NewHandler(challengerUsecase)
	getQuoteHandler := get_quote.NewHandler(
		challengerUsecase,
		quoterUsecase,
		powHeaderBuilderUsecase,
		e.Logger,
	)

	// Routes
	e.GET(internal.GetChallengeEndpoint, getChallengerHandler.GetChallenge)
	e.GET(internal.GetQuoteEndpoint, getQuoteHandler.GetQuote)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
