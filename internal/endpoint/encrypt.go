package endpoint

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) Encrypt(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil
	}

	service := a.serviceFactory.NewCryptService()

	content, err := service.Encrypt(body)
	if err != nil {
		return nil
	}

	return c.String(http.StatusOK, fmt.Sprintf("%x", content))
}
