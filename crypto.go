package jb

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// Creds are an example of stored Credentials
type Creds struct {
	Name     string
	Password Secret
}

// Secret stores a secret.
type Secret struct {
	Value []byte
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return cipherText, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], data)
	return cipherText, err
}

func decrypt(cipherText []byte, passphrase string) ([]byte, error) {
	if string(cipherText) == "ErrNotFound" {
		err := errors.New("value does not exist")
		return cipherText, err
	}
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return cipherText, err
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return cipherText, err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)
	return cipherText, err
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GetSHA returns the SHASUM for the given data.
func GetSHA(algorithm string, data []byte) string {
	match := true
	switch match {
	case algorithm == "sha1":
		return fmt.Sprintf("%x", sha1.Sum(data))
	case algorithm == "sha512":
		return fmt.Sprintf("%x", sha512.Sum512(data))
	}
	return fmt.Sprintf("%x", sha256.Sum256(data))
}
