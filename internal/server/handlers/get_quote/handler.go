package get_quote

import (
	"fmt"
	"net/http"

	"github.com/aed86/proof_of_work/internal"
	"github.com/aed86/proof_of_work/internal/pkg/challenger"
	"github.com/aed86/proof_of_work/internal/pkg/pow_header_builder"
	"github.com/aed86/proof_of_work/internal/pkg/quoter"
	"github.com/labstack/echo/v4"
)

type handler struct {
	challenger       challenger.Usecase
	quoter           quoter.Usecase
	powHeaderBuilder pow_header_builder.Usecase
}

func NewHandler(
	challenger challenger.Usecase,
	quoter quoter.Usecase,
	powHeaderBuilder pow_header_builder.Usecase,
) *handler {
	return &handler{
		challenger:       challenger,
		quoter:           quoter,
		powHeaderBuilder: powHeaderBuilder,
	}
}

func (h *handler) GetQuote(c echo.Context) error {
	powHash := c.Request().Header.Get(internal.PowHeaderName)
	if powHash == "" {
		return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("'%s' header is required", internal.PowHeaderName))
	}

	solution, err := h.powHeaderBuilder.Extract(powHash)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = h.challenger.Validate(*solution)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Challenge not solved")
	}

	quote, err := h.quoter.GetQuote()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, quote)
}
