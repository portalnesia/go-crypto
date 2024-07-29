/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package example

import (
	"fmt"
	"go.portalnesia.com/crypto"
)

func Password() {
	stringPassword := "Admin#1234"

	hashedPassword := crypto.HashPassword(stringPassword)
	fmt.Println("Hashed Password:", hashedPassword)

	match := crypto.ComparePassword(stringPassword, hashedPassword)
	if !match {
		fmt.Println("Passwords don't match")
	} else {
		fmt.Println("Passwords match")
	}
}
