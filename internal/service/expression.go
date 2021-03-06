package service

import (
	"encoding/base64"
	"strings"

	"github.com/sirupsen/logrus"
)

type parseExpression struct {
	source       map[string]interface{}
	cryptService CryptService
}

func newParseExpression(source map[string]interface{}, cryptService CryptService) *parseExpression {
	return &parseExpression{
		source:       source,
		cryptService: cryptService,
	}
}

func (e *parseExpression) eval() map[string]interface{} {
	result := map[string]interface{}{}

	for key, value := range e.source {
		switch value.(type) {
		case string:
			result[key] = e.decode(value.(string))
		default:
			result[key] = value
		}
	}

	return result
}

func (e *parseExpression) decode(source string) string {
	if strings.HasPrefix(source, "{cipher-base64}") {
		content := strings.ReplaceAll(source, "{cipher-base64}", "")
		content = strings.ReplaceAll(content, "\"", "")

		decrypted, err := e.cryptService.Decrypt(string(content))
		if err != nil {
			logrus.Error(err)
			return source
		}

		decoded, err := base64.StdEncoding.DecodeString(string(decrypted))
		if err != nil {
			logrus.Error(err)
			return ""
		}

		return string(decoded)
	}

	if strings.HasPrefix(source, "{cipher}") {
		content := strings.ReplaceAll(source, "{cipher}", "")
		content = strings.ReplaceAll(content, "\"", "")
		decoded, err := e.cryptService.Decrypt(content)
		if err != nil {
			logrus.Error(err)
			return source
		}
		return string(decoded)
	}

	return source
}
