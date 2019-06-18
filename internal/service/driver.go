package service

import (
	"log"
	"strings"

	"github.com/helderfarias/go-config-server/internal/domain"
)

type DriveNativeFactory interface {
	BuildProperySources() []domain.PropertySource
}

type emptyDriveNative struct {
}

func NewDriveNativeFactory(cloud domain.SpringCloudConfig, application, profile string) DriveNativeFactory {
	actives := strings.Split(cloud.Spring.Profiles.Active, ",")
	if isDriveInvalid(actives) {
		log.Println("Profile is not defined:", actives)
		return &emptyDriveNative{}
	}

	searchLocations := cloud.Spring.Cloud.Config.Server.Native["searchLocations"]
	if searchLocations == nil {
		return &emptyDriveNative{}
	}

	return &fileDriveNative{
		searchLocations: searchLocations.(string),
		application:     application,
		profile:         profile,
	}
}

func (e *emptyDriveNative) BuildProperySources() []domain.PropertySource {
	return []domain.PropertySource{}
}

func isDriveInvalid(actives []string) bool {
	for _, s := range actives {
		if strings.TrimSpace(s) == strings.TrimSpace("dev") ||
			strings.TrimSpace(s) == strings.TrimSpace("git") {
			return true
		}
	}

	return false
}
