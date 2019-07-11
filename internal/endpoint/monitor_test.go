package endpoint

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMonitorRefreshToken(t *testing.T) {
	event := `{"commits": [{"modified": ["accountservice.yml"]}],"name":"what is this?"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/monitor", strings.NewReader(event))
	req.Header.Set("X-Github-Event", "push")
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := domain.SpringCloudConfig{}
	cfg.Spring.Nats.Servers = "nats://localhost:5222"
	c.Set("cloudConfig", cfg)

	err := api.MonitorRefreshToken(c)

	assert.NoError(t, err)
	assert.Equal(t, "accountservice.yml", mockBroker.Message)
}
