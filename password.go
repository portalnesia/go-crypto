/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash password with bcrypt
func HashPassword(p string) string {
	password := []byte(p)

	hashed, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

// ComparePassword compare password with saved hash password
func ComparePassword(password string, hasedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
	return err == nil
}
