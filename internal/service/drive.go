package service

import (
	"strings"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/sirupsen/logrus"
)

type DriveNativeFactory interface {
	Build() *domain.BuildSource
}

type emptyDriveNative struct {
}

func newDriveNative(cfg domain.EnvConfig) DriveNativeFactory {
	actives := strings.Split(cfg.Cloud.Spring.Profiles.Active, ",")
	if isDriveInvalid(actives) {
		logrus.Warnf("Profile is not defined: %s", actives)
		return &emptyDriveNative{}
	}

	if len(actives) == 0 {
		return &gitDriveNative{
			source:       cfg.Cloud.Spring.Cloud.Config.Server.Git,
			application:  cfg.Application,
			profile:      cfg.Profile,
			label:        cfg.Label,
			cryptService: newCryptService(cfg),
		}
	}

	return createComposeDrive(actives, cfg)
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

func createComposeDrive(actives []string, cfg domain.EnvConfig) DriveNativeFactory {
	compose := &composeDriveNative{
		targets: []DriveNativeFactory{},
	}

	for i := len(actives) - 1; i >= 0; i-- {
		if strings.TrimSpace(actives[i]) == "native" {
			compose.Add(&fileDriveNative{
				index:        i,
				source:       cfg.Cloud.Spring.Cloud.Config.Server.Native,
				application:  cfg.Application,
				profile:      cfg.Profile,
				label:        cfg.Label,
				cryptService: newCryptService(cfg),
			})
		} else if strings.TrimSpace(actives[i]) == "git" {
			compose.Add(&gitDriveNative{
				index:        i,
				source:       cfg.Cloud.Spring.Cloud.Config.Server.Git,
				application:  cfg.Application,
				profile:      cfg.Profile,
				label:        cfg.Label,
				cryptService: newCryptService(cfg),
			})
		}
	}

	return compose
}
