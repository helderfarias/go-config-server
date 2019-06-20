package endpoint

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/helderfarias/go-config-server/internal/service"

	"github.com/labstack/echo"
)

func (*Api) Encrypt(c echo.Context) error {
	if c.Get("cloudConfig") == nil {
		return c.String(http.StatusOK, "")
	}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil
	}

	cloud := c.Get("cloudConfig").(domain.SpringCloudConfig)

	srv := service.NewCryptServiceFactory(domain.EnvConfig{
		Cloud: cloud,
	})

	content, err := srv.Encrypt(body)
	if err != nil {
		return nil
	}

	return c.String(http.StatusOK, fmt.Sprintf("%x", content))
}
