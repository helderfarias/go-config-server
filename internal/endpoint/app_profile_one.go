package endpoint

import (
	"net/http"

	"github.com/guregu/null"
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/helderfarias/go-config-server/internal/service"
	"github.com/labstack/echo"
)

func (*Api) ApplicationProfileOne(c echo.Context) error {
	if c.Get("cloudConfig") == nil {
		return c.JSON(http.StatusOK, domain.ProfileConfig{
			Name:     null.StringFrom(c.Param("application")),
			Profiles: []string{c.Param("profile")},
			Label:    null.StringFrom(c.Param("label")),
		})
	}

	cloud := c.Get("cloudConfig").(domain.SpringCloudConfig)

	driver := service.NewDriveNativeFactory(cloud, c.Param("application"), c.Param("profile"))

	sources := driver.BuildProperySources()

	return c.JSON(http.StatusOK, domain.ProfileConfig{
		Name:            null.StringFrom(c.Param("application")),
		Profiles:        []string{c.Param("profile")},
		Label:           null.StringFrom(c.Param("label")),
		PropertySources: sources,
	})
}
