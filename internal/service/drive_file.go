package service

import (
	"strings"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/sirupsen/logrus"
)

type fileDriveNative struct {
	source       map[string]interface{}
	application  string
	profile      string
	label        string
	index        int
	cryptService CryptService
}

func (e *fileDriveNative) Build() *domain.BuildSource {
	directory := e.source["searchLocations"].(string)

	resolver := newResolverFile(directory, e.application, e.profile)

	name, data, err := resolver.decode()
	if err != nil {
		logrus.Error(err)
		return &domain.BuildSource{}
	}

	source := map[string]interface{}{}
	for key, value := range data {
		switch value.(type) {
		case string:
			source[key] = e.eval(value.(string))
		default:
			source[key] = value
		}
	}

	return domain.NewBuildSource().
		AddProperty(domain.PropertySource{
			Name:   name,
			Source: source,
			Index:  e.index,
		})
}

func (e *fileDriveNative) eval(source string) string {
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
