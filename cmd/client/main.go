package main

import (
	"fmt"
	"log"

	"github.com/aed86/proof_of_work/internal/client"
	usecase2 "github.com/aed86/proof_of_work/internal/pkg/challenger/usecase"
	"github.com/aed86/proof_of_work/internal/pkg/pow_header_builder/usecase"
)

func main() {
	fmt.Println("Started")
	c := client.NewClient()

	fmt.Println("Getting challenge")
	challenge, err := c.GetChallenge()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Challenge received")
	fmt.Println(challenge)

	powHeaderBuilder := usecase.NewPowHeaderBuilder()
	challenger := usecase2.NewChallenger()

	solution := challenger.Solve(*challenge)
	fmt.Println(fmt.Sprintf("solution: %x", solution))

	powHeader := powHeaderBuilder.Build(*solution)
	fmt.Println(fmt.Sprintf("pow-header prepared: %s", powHeader))

	quote, err := c.GetQuote(powHeader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quote)
	fmt.Println("Finished")

}
