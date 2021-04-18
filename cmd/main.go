package main

import (
	"fmt"
	"net/http"

	env "github.com/dmitrymomot/go-env"
	"github.com/dmitrymomot/lfm"
	"github.com/jinzhu/configor"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Appliaction environment variables
var (
	appPort        = env.GetInt("APP_PORT", 8000)
	appBaseURL     = env.GetString("APP_BASE_URL", "http://localhost")
	debugMode      = env.GetBool("DEBUG_MODE", false)
	templateDir    = env.GetString("TEMPLATE_DIR", "./src/views")
	healthEndpoint = env.GetString("HEALTH_ENDPOINT", "/health")
	formConfigPath = env.GetString("FORM_CONFIG_PATH", "./form.config.yaml")
	staticFilesDir = env.GetString("STATIC_FILES_DIR", "./src/assets")
)

func main() {
	formConfig := &lfm.FormConfig{}
	configor.New(&configor.Config{Debug: debugMode}).Load(formConfig, formConfigPath)

	e := echo.New()
	e.Debug = debugMode
	e.HideBanner = !debugMode
	e.HTTPErrorHandler = lfm.HTTPErrorHandler

	e.Pre(middleware.NonWWWRedirect())
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.CSRF())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{appBaseURL},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodOptions,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	renderer, err := lfm.NewRenderer(templateDir)
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Renderer = renderer

	// Serve static files
	e.Static("/assets", staticFilesDir)

	// Set up endpoint for health check
	e.Any(healthEndpoint, func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appPort)))
}
