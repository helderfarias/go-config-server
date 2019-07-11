package endpoint

import (
	"net/http"

	"github.com/guregu/null"
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/labstack/echo/v4"
)

func (a *Api) ApplicationProfile(c echo.Context) error {
	driver := a.serviceFactory.NewDriveNative(domain.EnvConfig{
		Application: c.Param("application"),
		Profile:     c.Param("profile"),
		Label:       c.Param("label"),
	})

	build := driver.Build()

	return c.JSON(http.StatusOK, domain.ProfileConfig{
		Name:            null.StringFrom(c.Param("application")),
		Profiles:        []string{c.Param("profile")},
		Label:           null.StringFrom(c.Param("label")),
		Version:         null.StringFrom(build.Options["version"]),
		PropertySources: build.Properties,
	})
}
