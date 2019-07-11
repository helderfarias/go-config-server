package service

import (
	"github.com/helderfarias/go-config-server/internal/domain"
)

type composeDriveNative struct {
	targets []DriveNativeFactory
	source  map[string]interface{}
}

func (e *composeDriveNative) Add(newTarget DriveNativeFactory) {
	e.targets = append(e.targets, newTarget)
}

func (e *composeDriveNative) Build() *domain.BuildSource {
	data := domain.NewBuildSource()

	for _, d := range e.targets {
		result := d.Build()
		if len(result.Properties) > 0 {
			data.AddProperty(result.Properties[0])
		}

		if len(result.Options) > 0 {
			for k, v := range result.Options {
				data.AddOption(k, v)
			}
		}
	}

	return data
}
