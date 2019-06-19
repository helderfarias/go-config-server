package endpoint

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/helderfarias/go-config-server/internal/service"
	"github.com/labstack/echo"
)

func (*Api) Decrypt(c echo.Context) error {
	if c.Get("cloudConfig") == nil {
		return c.String(http.StatusOK, "")
	}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil
	}

	cloud := c.Get("cloudConfig").(domain.SpringCloudConfig)

	srv := service.NewCryptService(cloud.Encrypt.Key)

	content, err := srv.Decrypt(string(body))
	if err != nil {
		return nil
	}

	return c.String(http.StatusOK, fmt.Sprintf("%s", content))
}
