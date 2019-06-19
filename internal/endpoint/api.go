package endpoint

import (
	"github.com/labstack/echo"
)

type Api struct {
}

func Register(e *echo.Echo) {
	api := Api{}
	e.GET("/:application/:profile", api.ApplicationProfileOne)
	e.GET("/:application/:profile/:label", api.ApplicationProfileOne)
	e.GET("/:application_profile", api.ApplicationProfileTwo)
	e.POST("/encrypt", api.Encrypt)
	e.POST("/decrypt", api.Decrypt)
}
