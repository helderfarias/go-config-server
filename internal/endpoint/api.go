package endpoint

import (
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/helderfarias/go-config-server/internal/service"
	"github.com/labstack/echo/v4"
)

type Api struct {
	serviceFactory service.Factory
}

func Register(e *echo.Echo, cfg domain.SpringCloudConfig) {
	api := Api{
		serviceFactory: service.NewServiceFactory(cfg),
	}

	e.GET("/:application/:profile", api.ApplicationProfile)
	e.GET("/:application/:profile/:label", api.ApplicationProfile)
	e.POST("/encrypt", api.Encrypt)
	e.POST("/decrypt", api.Decrypt)
	e.POST("/monitor", api.MonitorRefreshToken)
}
