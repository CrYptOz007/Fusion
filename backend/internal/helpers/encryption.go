package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveKey(password string, salt []byte, keyLen int) []byte {
	return pbkdf2.Key([]byte(password), salt, 4096, keyLen, sha256.New)
}

func Encrypt(text, password, salt string) (string, error) {

	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", err
	}

	key := DeriveKey(password, saltBytes, 16)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)

	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, key)
	stream.XORKeyStream(ciphertext, plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encryptedText, password, salt string) (string, error) {
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", err
	}

	key := DeriveKey(password, saltBytes, 16)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, _ := base64.URLEncoding.DecodeString(encryptedText)

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, key)
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

func GenerateRandomKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
			log.Fatalf("Failed to generate random key: %s", err)
	}
	return base64.StdEncoding.EncodeToString(key)
}