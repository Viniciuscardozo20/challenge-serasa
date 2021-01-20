package cryptoModule

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
)

const nonce = "noncekey0001"

func Encrypt(data []byte, passphrase string) (string, error) {
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return "", errors.Wrap(err, "failed to create Cipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "failed to create GCM")
	}
	ciphertext := gcm.Seal(nil, []byte(nonce), data, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(data string, passphrase string) (string, error) {
	dataDecoded, err := hex.DecodeString(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode string")
	}
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return "", errors.Wrap(err, "failed to create Cipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "failed to create GCM")
	}
	plaintext, err := gcm.Open(nil, []byte(nonce), dataDecoded, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt")
	}
	return string(plaintext), nil
}
