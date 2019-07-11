package endpoint

import (
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/helderfarias/go-config-server/internal/service"
)

type mockServiceFactory struct {
	mockBroker      *mockMessageBroker
	mockDriveNative *mockDriveNative
	mockCrypt       *mockCryptService
}

type mockMessageBroker struct {
	Message string
}

type mockDriveNative struct {
}

type mockCryptService struct {
}

func (m *mockMessageBroker) Publish(msg string) error {
	m.Message = msg
	return nil
}

func (m *mockServiceFactory) NewDriveNative(cfg domain.EnvConfig) service.DriveNativeFactory {
	return m.mockDriveNative
}

func (m *mockServiceFactory) NewCryptService() service.CryptService {
	return m.mockCrypt
}

func (m *mockServiceFactory) NewMessageBroker() service.MessageBroker {
	return m.mockBroker
}

func (m *mockDriveNative) Build() *domain.BuildSource {
	return &domain.BuildSource{}
}

func (m *mockCryptService) Encrypt(data []byte) ([]byte, error) {
	return []byte("encrypted"), nil
}

func (m *mockCryptService) Decrypt(source string) ([]byte, error) {
	return []byte("decrypted"), nil
}

var mockBroker *mockMessageBroker
var mockFactory *mockServiceFactory
var mockDrive *mockDriveNative
var mockCrypt *mockCryptService
var api Api

func init() {
	mockCrypt = &mockCryptService{}
	mockDrive = &mockDriveNative{}
	mockBroker = &mockMessageBroker{}
	mockFactory = &mockServiceFactory{
		mockBroker:      mockBroker,
		mockDriveNative: mockDrive,
		mockCrypt:       mockCrypt,
	}
	api = Api{serviceFactory: mockFactory}
}
