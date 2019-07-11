package endpoint

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) Decrypt(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil
	}

	service := a.serviceFactory.NewCryptService()

	content, err := service.Decrypt(string(body))
	if err != nil {
		return nil
	}

	return c.String(http.StatusOK, string(content))
}
