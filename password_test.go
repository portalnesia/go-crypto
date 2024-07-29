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

var (
	pass   = "passwordfromtesting"
	hashed = HashPassword(pass)
)

func TestCompareHashedPassword(t *testing.T) {
	ok := ComparePassword(pass, hashed)
	if !ok {
		t.Errorf("Password not match")
	}
}

func TestNodeJsComparePassword(t *testing.T) {
	godotenv.Load(".env")
	fmt.Println("TES", os.Getenv("NODEJS_PASS"), os.Getenv("NODEJS_HASHED"))
	ok := ComparePassword(os.Getenv("NODEJS_PASS"), os.Getenv("NODEJS_HASHED"))
	if !ok {
		t.Errorf("Password not match")
	}
}
