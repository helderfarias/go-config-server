package service

import (
	"sort"
	"sync"

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
	var waitgroup sync.WaitGroup
	var semaforo sync.Mutex

	data := domain.NewBuildSource()

	for _, inner := range e.targets {
		waitgroup.Add(1)
		go func(waitgroup *sync.WaitGroup, d DriveNativeFactory) {
			result := d.Build()

			if len(result.Properties) > 0 {
				semaforo.Lock()
				data.AddProperty(domain.PropertySource{
					Name:   result.Properties[0].Name,
					Index:  result.Properties[0].Index,
					Source: result.Properties[0].Source,
				})
				semaforo.Unlock()
			}

			if len(result.Options) > 0 {
				semaforo.Lock()
				for k, v := range result.Options {
					data.AddOption(k, v)
				}
				semaforo.Unlock()
			}

			waitgroup.Done()
		}(&waitgroup, inner)
	}

	waitgroup.Wait()

	sort.SliceStable(data.Properties, func(a, b int) bool {
		return data.Properties[a].Index > data.Properties[b].Index
	})

	return data
}
