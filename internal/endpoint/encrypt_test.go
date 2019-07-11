package endpoint

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestShouldEncrypt(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/encrypt", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := api.Encrypt(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
