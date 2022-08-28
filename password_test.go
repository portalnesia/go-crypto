package crypto

import (
	"fmt"
	"os"
	"testing"
)

var (
	PASS   = "passwordfromtesting"
	HASHED string
)

func TestHashPassword(t *testing.T) {
	HASHED = HashPassword(PASS)
	fmt.Println(HASHED)
}

func TestCompareHashedPassword(t *testing.T) {
	ok := ComparePassword(PASS, HASHED)
	if !ok {
		t.Errorf("Password not match")
	}
}

func TestNodeJsComparePassword(t *testing.T) {
	ok := ComparePassword(os.Getenv("NODEJS_PASS"), os.Getenv("NODEJS_HASHED"))
	if !ok {
		t.Errorf("Password not match")
	}
}
