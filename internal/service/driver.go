package service

import (
	"log"
	"strings"

	"github.com/helderfarias/go-config-server/internal/domain"
)

type DriveNativeFactory interface {
	Build() *domain.BuildSource
}

type emptyDriveNative struct {
}

func NewDriveNativeFactory(cfg domain.EnvConfig) DriveNativeFactory {
	actives := strings.Split(cfg.Cloud.Spring.Profiles.Active, ",")
	if isDriveInvalid(actives) {
		log.Println("Profile is not defined:", actives)
		return &emptyDriveNative{}
	}

	cryptService := NewCryptServiceFactory(cfg)

	if actives[0] == "native" {
		return &fileDriveNative{
			source:       cfg.Cloud.Spring.Cloud.Config.Server.Native,
			application:  cfg.Application,
			profile:      cfg.Profile,
			label:        cfg.Label,
			cryptService: cryptService,
		}
	}

	if actives[0] == "git" {
		return &gitDriveNative{
			source:       cfg.Cloud.Spring.Cloud.Config.Server.Git,
			application:  cfg.Application,
			profile:      cfg.Profile,
			label:        cfg.Label,
			cryptService: cryptService,
		}
	}

	return &emptyDriveNative{}
}

func (e *emptyDriveNative) Build() *domain.BuildSource {
	return domain.NewBuildSource()
}

func isDriveInvalid(actives []string) bool {
	for _, s := range actives {
		if strings.TrimSpace(s) == strings.TrimSpace("native") ||
			strings.TrimSpace(s) == strings.TrimSpace("git") {
			return false
		}
	}

	return true
}
