package service

import (
	"log"

	"github.com/helderfarias/go-config-server/internal/domain"
)

type fileDriveNative struct {
	source      map[string]interface{}
	application string
	profile     string
	label       string
}

func (e *fileDriveNative) Build() *domain.BuildSource {
	directory := e.source["searchLocations"].(string)

	resolver := newResolverFile(directory, e.application, e.profile)

	name, source, err := resolver.decode()
	if err != nil {
		log.Println(err)
		return &domain.BuildSource{}
	}

	return domain.NewBuildSource().
		AddProperty(domain.PropertySource{
			Name:   name,
			Source: source,
		})
}
