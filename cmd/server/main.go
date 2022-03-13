package main

import (
	"github.com/aed86/proof_of_work/internal/server/handlers/get_challenge"
	"github.com/aed86/proof_of_work/internal/server/handlers/get_quote"
	"github.com/aed86/proof_of_work/internal/server/pkg/challenger/usecase"
	quoter_usecase "github.com/aed86/proof_of_work/internal/server/pkg/quoter/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	challengerUsecase := usecase.NewChallenger()
	quoterUsecase := quoter_usecase.NewQuoter()

	getChallengerHandler := get_challenge.NewHandler(challengerUsecase)
	getQuoteHandler := get_quote.NewHandler(
		challengerUsecase,
		quoterUsecase,
	)
	//errorHandler := error2.NewHandler()

	// Routes
	e.GET("/getChallenge", getChallengerHandler.GetChallenge)
	e.GET("/getQuote", getQuoteHandler.GetQuote)

	//e.HTTPErrorHandler = errorHandler.Handler

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
