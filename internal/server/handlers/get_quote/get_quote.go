package get_quote

import (
	"fmt"
	"net/http"

	"github.com/aed86/proof_of_work/internal/server/pkg/challenger"
	"github.com/aed86/proof_of_work/internal/server/pkg/quoter"
	"github.com/labstack/echo/v4"
)

type handler struct {
	challenger challenger.Usecase
	quoter     quoter.Usecase
}

const PowHeaderName = "X-POW-Auth"

func NewHandler(
	challenger challenger.Usecase,
	quoter quoter.Usecase,
) *handler {
	return &handler{
		challenger: challenger,
		quoter:     quoter,
	}
}

func (h *handler) GetQuote(c echo.Context) error {
	powHash := c.Request().Header.Get(PowHeaderName)
	if powHash == "" {
		return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("'%s' header is required", PowHeaderName))
	}

	solveData, err := h.challenger.Extract(powHash)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = h.challenger.Validate(*solveData)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Challenge not solved")
	}

	quote, err := h.quoter.GetQuote()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, quote)
}
