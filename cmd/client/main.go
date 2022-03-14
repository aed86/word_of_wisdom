package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/aed86/proof_of_work/cmd"
	"github.com/aed86/proof_of_work/internal/client"
	challenger_usecase "github.com/aed86/proof_of_work/internal/pkg/challenger/usecase"
	"github.com/aed86/proof_of_work/internal/pkg/pow_header_builder/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	var (
		c      = client.NewClient(cmd.GetAddressFromEnv())
		e      = echo.New()
		logger = e.Logger
	)

	challenge, err := c.GetChallenge()
	if err != nil {
		logger.Fatal(err)
	}

	err = prettyPrint("Got challenge:", challenge)
	if err != nil {
		logger.Fatal(err)
	}

	powHeaderBuilder := usecase.NewPowHeaderBuilder()
	challenger := challenger_usecase.NewChallenger(logger)

	solution := challenger.Solve(*challenge)
	err = prettyPrint("Solution: ", solution)
	if err != nil {
		logger.Fatal(err)
	}

	powHeader := powHeaderBuilder.Build(*solution)
	quote, err := c.GetQuote(powHeader)
	if err != nil {
		logger.Fatal(err)
	}

	err = prettyPrint("Random quote: ", quote)
	if err != nil {
		logger.Fatal(err)
	}
}

func prettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func prettyPrint(topic string, json interface{}) error {
	var buffer bytes.Buffer
	err := prettyEncode(json, &buffer)
	if err != nil {
		return err
	}

	fmt.Println(topic)
	fmt.Println(buffer.String())

	return nil
}
