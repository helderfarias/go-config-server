package middleware

import (
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/labstack/echo/v4"
)

func AuthApiKey(cfg domain.SpringCloudConfig) echo.MiddlewareFunc {
	return KeyAuthWithConfig(
		KeyAuthConfig{
			Skipper: func(echo.Context) bool {
				return !cfg.Security.APIKey.Enabled
			},
			KeyLookup: cfg.Security.APIKey.KeyLookup,
			Validator: func(key string, c echo.Context) (bool, error) {
				return key == cfg.Security.APIKey.Token, nil
			},
		},
	)
}
