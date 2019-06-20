package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type cryptServiceDefault struct {
	masterKey string
}

func (e *cryptServiceDefault) Encrypt(data []byte) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(e.createHash(e.masterKey)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, err
}

func (e *cryptServiceDefault) Decrypt(source string) ([]byte, error) {
	data, err := hex.DecodeString(source)
	if err != nil {
		return nil, err
	}

	key := []byte(e.createHash(e.masterKey))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, err
}

func (e *cryptServiceDefault) createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
