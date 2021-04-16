package lfm

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPErrorHandler
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var desc interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		desc = he.Message
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.Render(code, errorPage, echo.Map{
		"error_code":        code,
		"error_message":     http.StatusText(code),
		"error_description": desc,
	}); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
