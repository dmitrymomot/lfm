package lfm

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PageHandler renders html pages,
// based on html files from templates directory
func PageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, c.Param("page"), nil)
}

// ErrorPageHandler renders error page
func ErrorPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, c.Param("page"), nil)
}
