package service

import (
	"fmt"
	"testing"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestShouldDecryptWithParseExpressionUsingCipherBase64(t *testing.T) {
	global := domain.SpringCloudConfig{}
	global.Encrypt.Key = "teste01"
	plainCert := `
-----BEGIN PRIVATE KEY-----
GfgfgCCqGSM49AAgEGCCqGSM49AKKK9AGfgfgCCqGSM49AAgEGCCqGSM49AKKK9A
GfgfgCCqGSM49AAgEGCCqGSM49AKKK9AGfgfgCCqGSM49AAgEGCCqGSM49AKKK9A
GfgfgCCqGSM49AAgEGCCqG
-----END PRIVATE KEY-----
	`
	b64Cert := "Ci0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLQpHZmdmZ0NDcUdTTTQ5QUFnRUdDQ3FHU000OUFLS0s5QUdmZ2ZnQ0NxR1NNNDlBQWdFR0NDcUdTTTQ5QUtLSzlBCkdmZ2ZnQ0NxR1NNNDlBQWdFR0NDcUdTTTQ5QUtLSzlBR2ZnZmdDQ3FHU000OUFBZ0VHQ0NxR1NNNDlBS0tLOUEKR2ZnZmdDQ3FHU000OUFBZ0VHQ0NxRwotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCgk="

	service := newCryptService(domain.EnvConfig{Cloud: global})

	encoded, _ := service.Encrypt([]byte(string(b64Cert)))

	parse := newParseExpression(map[string]interface{}{
		"cert": fmt.Sprintf("{cipher-base64}%x", encoded),
	}, service)

	assert.Equal(t, plainCert, parse.eval()["cert"])
}

func TestShouldDecryptWithParseExpressionUsingCipher(t *testing.T) {
	global := domain.SpringCloudConfig{}
	global.Encrypt.Key = "teste01"
	plainCert := "hi"

	service := newCryptService(domain.EnvConfig{Cloud: global})

	encoded, _ := service.Encrypt([]byte(string(plainCert)))

	parse := newParseExpression(map[string]interface{}{
		"cert": fmt.Sprintf("{cipher}%x", encoded),
	}, service)

	assert.Equal(t, plainCert, parse.eval()["cert"])
}

func TestNotShouldDecryptWithParseExpressionInvalid(t *testing.T) {
	global := domain.SpringCloudConfig{}
	global.Encrypt.Key = "teste01"
	plainCert := "hi"

	service := newCryptService(domain.EnvConfig{Cloud: global})

	encoded, _ := service.Encrypt([]byte(string(plainCert)))

	parse := newParseExpression(map[string]interface{}{
		"cert": fmt.Sprintf("%x", encoded),
	}, service)

	assert.Equal(t, fmt.Sprintf("%x", encoded), parse.eval()["cert"])
}
