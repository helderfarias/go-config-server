package endpoint

import (
	"net/http"
	"strings"

	"github.com/guregu/null"
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/helderfarias/go-config-server/internal/service"
	"github.com/labstack/echo"
)

func (*Api) ApplicationProfileTwo(c echo.Context) error {
	applicationProfile := strings.Split(strings.TrimSuffix(c.Param("application_profile"), ".yml"), "-")

	if c.Get("cloudConfig") == nil {
		return c.JSON(http.StatusOK, domain.ProfileConfig{
			Name:     null.StringFrom(applicationProfile[0]),
			Profiles: []string{applicationProfile[1]},
		})
	}

	if len(applicationProfile) != 2 {
		return c.JSON(http.StatusOK, domain.ProfileConfig{
			Name:     null.StringFrom(applicationProfile[0]),
			Profiles: []string{applicationProfile[1]},
		})
	}

	cloud := c.Get("cloudConfig").(domain.SpringCloudConfig)

	driver := service.NewDriveNativeFactory(cloud, applicationProfile[0], applicationProfile[1])

	sources := driver.BuildProperySources()

	return c.JSON(http.StatusOK, domain.ProfileConfig{
		Name:            null.StringFrom(applicationProfile[0]),
		Profiles:        []string{applicationProfile[1]},
		PropertySources: sources,
	})
}
