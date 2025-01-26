/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package examples

import (
	"fmt"
	"go.portalnesia.com/crypto"
)

func Crypto() {
	c := crypto.New("YOUR_SECRET_KEY")

	text := "Hello World"

	// Encrypted Text
	encrypted, err := c.Encrypt(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted Text: %s", encrypted)

	// Decrypted Text
	decrypted, err := c.Decrypt(encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted Text: %s", decrypted)
}
