package error

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type errorHandler struct {
}

func NewHandler() *errorHandler {
	return &errorHandler{}
}

func (h *errorHandler) Handler(err error, c echo.Context) {
	hErr, ok := err.(*echo.HTTPError)
	if ok {
		if hErr.Internal != nil {
			if herr, ok := hErr.Internal.(*echo.HTTPError); ok {
				hErr = herr
			}
		}
	} else {
		hErr = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	code := hErr.Code
	message := hErr.Message
	if _, ok := hErr.Message.(string); ok {
		message = map[string]interface{}{"message": err.Error()}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(hErr.Code)
		} else {
			err = c.JSON(code, message)
		}

		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
