package endpoint

import (
	"net/http"

	"github.com/dimiro1/health"
	"github.com/labstack/echo/v4"
)

func (a *Api) Health(c echo.Context) error {
	handler := health.NewHandler()
	check := handler.CompositeChecker.Check()
	return c.JSON(http.StatusOK, check)
}
