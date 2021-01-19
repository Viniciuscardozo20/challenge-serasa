package cryptoModule

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

func createhash(key string) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(key))
	if err != nil {
		return "", errors.Wrap(err, "failed to create Hash")
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func Encrypt(data []byte, passphrase string) (string, error) {
	hash, err := createhash(passphrase)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(hash))
	if err != nil {
		return "", errors.Wrap(err, "failed to create Cipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "failed to create GCM")
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", errors.Wrap(err, "failed to read")
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(data string, passphrase string) (string, error) {
	decoded, err := hex.DecodeString(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode string")
	}
	hash, err := createhash(passphrase)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(hash))
	if err != nil {
		return "", errors.Wrap(err, "failed to create Cipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "failed to create GCM")
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := decoded[:nonceSize], decoded[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt")
	}
	return string(plaintext), nil
}
