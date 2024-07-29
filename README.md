[![Go Reference](https://pkg.go.dev/badge/go.portalnesia.com/crypto.svg)](https://pkg.go.dev/go.portalnesia.com/crypto) ![Unit Testing](https://github.com/portalnesia/go-crypto/actions/workflows/crypto_test.yml/badge.svg)

# Crypto

Crypto package for Internal Portalnesia

## Install

```bash
go get go.portalnesia.com/crypto
```

## Usage

### Crypto

```go
package main

import (
  "go.portalnesia.com/crypto"
  "fmt"
)

func main() {
	c := crypto.New("YOUR_SECRET_KEY")

	text := "Hello World"

	// Encrypted Text
	encrypted,err := c.Encrypt(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted Text: %s",encrypted)

	// Decrypted Text
	decrypted,err := c.Decrypt(encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted Text: %s",decrypted)
}
```

### Password

```go
package main

import (
	"fmt"
	"go.portalnesia.com/crypto"
)

func main() {
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

```

## Go References
[pkg.go.dev/go.portalnesia.com/crypto](https://pkg.go.dev/go.portalnesia.com/crypto)