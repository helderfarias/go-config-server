package middleware

import (
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func AuthApiKey(cfg domain.SpringCloudConfig) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(
		middleware.KeyAuthConfig{
			Skipper: func(echo.Context) bool {
				return !cfg.Security.APIKey.Enabled
			},
			KeyLookup: "query:apikey",
			Validator: func(key string, c echo.Context) (bool, error) {
				return key == cfg.Security.APIKey.Token, nil
			},
		},
	)
}
