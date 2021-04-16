package main

import (
	"fmt"
	"net/http"
	"time"

	env "github.com/dmitrymomot/go-env"
	"github.com/dmitrymomot/lfm"
	"github.com/jinzhu/configor"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Appliaction environment variables
var (
	appPort        = env.GetInt("APP_PORT", 8080)
	appBaseURL     = env.GetString("APP_BASE_URL", "http://localhost")
	debugMode      = env.GetBool("DEBUG_MODE", false)
	templateDir    = env.GetString("TEMPLATE_DIR", "./views")
	requestTimeout = env.GetDuration("REQUEST_TIMEOUT", 10*time.Second)
	healthEndpoint = env.GetString("HEALTH_ENDPOINT", "/health")
	formConfigPath = env.GetString("FORM_CONFIG_PATH", "./form.yml")
)

func main() {
	formConfig := &lfm.FormConfig{}
	configor.New(&configor.Config{Debug: debugMode}).Load(formConfig, formConfigPath)

	e := echo.New()
	e.Debug = debugMode
	e.HideBanner = !debugMode

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
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: requestTimeout}))

	renderer, err := lfm.NewRenderer(templateDir)
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Renderer = renderer

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
