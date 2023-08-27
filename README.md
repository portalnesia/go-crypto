[![Go Reference](https://pkg.go.dev/badge/go.portalnesia.com/crypto.svg)](https://pkg.go.dev/go.portalnesia.com/crypto) ![Unit Testing](https://github.com/portalnesia/go-crypto/actions/workflows/crypto_test.yml/badge.svg)

# Crypto

Crypto package for Internal Portalnesia

## Install

```bash
go get go.portalnesia.com/crypto
```

## Usage

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
  encrypted := c.Encrypt(text)
  fmt.Printf("Encrypted Text: %s",encrypted)

  // Decrypted Text
  decrypted := c.Decrypt(encrypted)
  fmt.Printf("Decrypted Text: %s",decrypted)
}
```

## Go References
[pkg.go.dev/go.portalnesia.com/crypto](https://pkg.go.dev/go.portalnesia.com/crypto)