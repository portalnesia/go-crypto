/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package crypto

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func GetConfig() (string, string) {
	const TEXT = "Hello world, this is me from go test"
	godotenv.Load(".env")
	return TEXT, os.Getenv("SECRET")
}

func TestCrypto(t *testing.T) {
	TEXT, KEY := GetConfig()
	crypto := New(KEY)
	encrypted, err := crypto.Encrypt(TEXT)

	if err != nil {
		t.Error(err)
	}

	decrypted, err := crypto.Decrypt(encrypted)

	fmt.Printf("Decrypted: %s\n", decrypted)

	if err != nil {
		t.Error(err)
	}

	if decrypted != TEXT {
		t.Errorf("Decrypted and Text not same")
	}
}

func TestCrypto_DecryptError(t *testing.T) {
	_, KEY := GetConfig()
	crypto := New(KEY)
	decrypted, err := crypto.Decrypt("some-text-123")
	if err == nil {
		t.Errorf("decrypt should fail")
	}
	if decrypted != "" {
		t.Errorf("decrypt should be empty")
	}
}

func TestFromNodeJs(t *testing.T) {
	TEXT, KEY := GetConfig()

	const encryptedFromNodeJs = "af215d824448b4b21d05a90768ae289a:01d3612196d88e088ffe22695078fe0cfb27af565c438d7c2680c14fd84b78645d28eef1c32f4e5e723f9e3e9e3a6811"

	crypto := New(KEY)

	decrypted, err := crypto.Decrypt(encryptedFromNodeJs)

	if err != nil {
		t.Error(err)
	}

	if decrypted != TEXT {
		t.Errorf("Decrypted and Text not same")
	}
}
