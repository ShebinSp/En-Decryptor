package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strconv"
)

func generateRandomKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// initializing for test purpose
var id int

func generateID() int {
	id++
	return id
}

func (i *imageData) encryptAES(plainData []byte) error {

	// Generate AES key
	aesKey, err := generateRandomKey()
	if err != nil {
		return err
	}
	i.aecKey = aesKey

	// Hashing the aesKey
	_ = i.hashKey()

	// Initialize AES cipher
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nounce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nounce)
	if err != nil {
		return err
	}

	cipherData := gcm.Seal(nounce, nounce, plainData, nil)

	i.data = cipherData

	return nil
}

func (i *imageData) decryptAES() ([]byte, error) {

	// Initialize AES cipher
	block, err := aes.NewCipher(i.aecKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nounceSize := gcm.NonceSize()

	nounce, cipherText := i.data[:nounceSize], i.data[nounceSize:]

	plainData, err := gcm.Open(nil, nounce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainData, nil
}

func (i *imageData) hashKey() string {
	hash := sha256.Sum256(i.aecKey)

	keyHash := hex.EncodeToString(hash[:])
	id := strconv.Itoa(i.id)
	keyHash += id

	i.hash = keyHash

	return keyHash
}