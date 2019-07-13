package service

import (
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

	resolver := newResolveFile(directory, e.application, e.profile)

	name, data, err := resolver.decode()
	if err != nil {
		logrus.Error(err)
		return &domain.BuildSource{}
	}

	parse := newParseExpression(data, e.cryptService)
	source := parse.eval()
	return domain.NewBuildSource().
		AddProperty(domain.PropertySource{
			Name:   name,
			Source: source,
			Index:  e.index,
		})
}
