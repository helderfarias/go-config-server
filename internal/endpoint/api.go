package endpoint

import (
	"github.com/labstack/echo"
)

type Api struct {
}

func Register(e *echo.Echo) {
	api := Api{}
	e.GET("/:application/:profile", api.ApplicationProfile)
	e.GET("/:application/:profile/:label", api.ApplicationProfile)
	e.POST("/encrypt", api.Encrypt)
	e.POST("/decrypt", api.Decrypt)
}
