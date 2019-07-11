package middleware

import (
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/labstack/echo/v4"
)

func SetCloudConfig(cloudConfig domain.SpringCloudConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("cloudConfig", cloudConfig)
			return next(c)
		}
	}
}
