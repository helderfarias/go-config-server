package service

import (
	"errors"

	"github.com/helderfarias/go-config-server/internal/domain"
)

type CryptService interface {
	Encrypt(data []byte) ([]byte, error)

	Decrypt(source string) ([]byte, error)
}

type cryptServiceEmpty struct {
}

func NewCryptServiceFactory(cfg domain.EnvConfig) CryptService {
	if cfg.Cloud.Encrypt.Key != "" {
		return &cryptServiceDefault{
			masterKey: cfg.Cloud.Encrypt.Key,
		}
	}

	return &cryptServiceEmpty{}
}

func (e *cryptServiceEmpty) Encrypt(data []byte) ([]byte, error) {
	return nil, errors.New("Encrypt is not enabled")
}

func (e *cryptServiceEmpty) Decrypt(source string) ([]byte, error) {
	return nil, errors.New("Encrypt is not enabled")
}
