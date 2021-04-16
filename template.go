package lfm

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
)

// NewRenderer is a factory function,
// return a new instanccec of the echo.Renderer interface implementation
func NewRenderer(templatesDir string) (echo.Renderer, error) {
	dir := fmt.Sprintf("%s/*.html", strings.TrimRight(templatesDir, "/"))
	tpl, err := template.ParseGlob(dir)
	if err != nil {
		return nil, fmt.Errorf("could not parse template directory %s: %w", dir, err)
	}
	return &renderer{templates: tpl}, nil
}

type renderer struct {
	templates *template.Template
}

func (r *renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}
