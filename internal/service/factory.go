package service

import "github.com/helderfarias/go-config-server/internal/domain"

type Factory interface {
	NewDriveNative(cfg domain.EnvConfig) DriveNativeFactory

	NewCryptService() CryptService

	NewMessageBroker() MessageBroker
}

type serviceFactory struct {
	springCloudConfig domain.SpringCloudConfig
}

func NewServiceFactory(cfg domain.SpringCloudConfig) Factory {
	return &serviceFactory{springCloudConfig: cfg}
}

func (s *serviceFactory) NewDriveNative(cfg domain.EnvConfig) DriveNativeFactory {
	return newDriveNative(domain.EnvConfig{
		Cloud:       s.springCloudConfig,
		Application: cfg.Application,
		Profile:     cfg.Profile,
		Label:       cfg.Label,
	})
}

func (s *serviceFactory) NewCryptService() CryptService {
	return newCryptService(domain.EnvConfig{
		Cloud: s.springCloudConfig,
	})
}

func (s *serviceFactory) NewMessageBroker() MessageBroker {
	return newFactoryMessageBroker(s.springCloudConfig)
}
