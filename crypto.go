/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/mergermarket/go-pkcs7"
)

type Crypto struct {
	key []byte
}

// New initialize Crypto instance
func New(secret string) Crypto {
	key := []byte(secret)
	return Crypto{key: key}
}

// Encrypt encrypt data
func (c Crypto) Encrypt(data string) (string, error) {
	plainText := []byte(data)
	plainText, err := pkcs7.Pad(plainText, aes.BlockSize)

	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}

	if len(plainText)%aes.BlockSize != 0 {
		err = fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	encryptedText := hex.EncodeToString(cipherText)
	encrypted := encryptedText[:32] + ":" + encryptedText[32:]

	return encrypted, nil
}

// Decrypt decrypt data
func (c Crypto) Decrypt(encrypted string) (string, error) {
	if len(encrypted) <= 0 {
		return "", nil
	}

	split := strings.Split(encrypted, ":")

	encrypted = split[0] + split[1]

	cipherText, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		return "", fmt.Errorf("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, err = pkcs7.Unpad(cipherText, aes.BlockSize)
	if err != nil {
		return "", err
	}

	return string(cipherText), nil
}
