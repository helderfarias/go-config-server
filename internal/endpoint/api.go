package endpoint

import (
	"github.com/helderfarias/go-config-server/internal/domain"
	mwi "github.com/helderfarias/go-config-server/internal/middleware"
	"github.com/helderfarias/go-config-server/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Api struct {
	serviceFactory service.Factory
}

func NewApi() {
	return
}

func Register(e *echo.Echo, cfg domain.SpringCloudConfig) {
	api := Api{
		serviceFactory: service.NewServiceFactory(cfg),
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.GET("/health", api.Health)

	sec := e.Group("")
	sec.Use(mwi.AuthBasic(cfg))
	sec.Use(mwi.AuthApiKey(cfg))
	sec.GET("/:application/:profile", api.ApplicationProfile)
	sec.GET("/:application/:profile/:label", api.ApplicationProfile)
	sec.POST("/encrypt", api.Encrypt)
	sec.POST("/decrypt", api.Decrypt)
	sec.POST("/monitor", api.MonitorRefreshToken)
}
