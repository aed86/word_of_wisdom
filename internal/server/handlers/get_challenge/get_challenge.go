package get_challenge

import (
	"net/http"

	"github.com/aed86/proof_of_work/internal/server/pkg/challenger"
	"github.com/labstack/echo/v4"
)

type handler struct {
	challenger challenger.Usecase
}

func NewHandler(challenger challenger.Usecase) *handler {
	return &handler{
		challenger: challenger,
	}
}

func (h *handler) GetChallenge(c echo.Context) error {
	challenge, err := h.challenger.GetChallenge()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, challenge)
}
