// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

func createHash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

// Encrypt plain text using the key
func Encrypt(plaintext string, salt string) ([]byte, error) {
	bytetext := []byte(plaintext)
	key := createHash(salt)
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, bytetext, nil), nil
}

// Decrypt cipher text using the key ...
func Decrypt(ciphertext string, salt string) ([]byte, error) {
	bytetext := []byte(ciphertext)
	key := createHash(salt)
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(bytetext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, bytetext := bytetext[:nonceSize], bytetext[nonceSize:]
	return gcm.Open(nil, nonce, bytetext, nil)
}

//DecodeAndDecrypt decode given ciphertext into base64 and apply decryption
func DecodeAndDecrypt(ciphertext string, salt string) ([]byte, error) {
	bytetext := DecodeBase64String(ciphertext)
	return Decrypt(string(bytetext), salt)
}

//DecodeBase64String decode base64 string
func DecodeBase64String(encodedStr string) []byte {
	decStr, err := base64.RawStdEncoding.DecodeString(encodedStr)
	if nil != err {
		return nil
	}
	return decStr
}

//EncodeBase64String encode to string
func EncodeBase64String(encodedStr []byte) string {
	return base64.RawStdEncoding.EncodeToString(encodedStr)
}
