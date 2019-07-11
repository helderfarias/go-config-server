package middleware

import (
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthBasic(cfg domain.SpringCloudConfig) echo.MiddlewareFunc {
	return middleware.BasicAuthWithConfig(
		middleware.BasicAuthConfig{
			Skipper: func(echo.Context) bool {
				return !cfg.Security.Basic.Enabled
			},
			Validator: func(username, password string, c echo.Context) (bool, error) {
				if username == cfg.Security.Basic.User &&
					password == cfg.Security.Basic.Password {
					return true, nil
				}
				return false, nil
			},
		},
	)
}
