package endpoint

import (
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/helderfarias/go-config-server/internal/service"
	"github.com/labstack/echo"
)

type Api struct {
	messageBroker service.MessageBroker
}

func Register(e *echo.Echo, cfg domain.SpringCloudConfig) {
	api := Api{
		messageBroker: service.NewFactoryMessageBroker(cfg),
	}

	e.GET("/:application/:profile", api.ApplicationProfile)
	e.GET("/:application/:profile/:label", api.ApplicationProfile)
	e.POST("/encrypt", api.Encrypt)
	e.POST("/decrypt", api.Decrypt)
	e.POST("/monitor", api.MonitorRefreshToken)
}
