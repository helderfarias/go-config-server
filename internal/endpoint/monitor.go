package endpoint

import (
	"net/http"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (a *Api) MonitorRefreshToken(c echo.Context) error {
	var data domain.GithubEvent
	if err := c.Bind(&data); err != nil {
		logrus.Error(err)
		return err
	}

	if c.Request().Header.Get("X-Github-Event") != "push" {
		logrus.Error("Evento esperado:", c.Request().Header.Get("X-Github-Event"))
		return c.JSON(http.StatusBadRequest, "Evento não permitido")
	}

	target := "empty"
	for _, resource := range data.Commits {
		if len(resource.Modified) > 0 {
			target = resource.Modified[0]
		}
	}

	broker := a.serviceFactory.NewMessageBroker()

	if err := broker.Publish(target); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
