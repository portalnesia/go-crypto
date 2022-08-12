# Go-Crypto

Crypto package for Internal Portalnesia

## Install

```bash
go get github.com/portalnesia/go-crypto
```

## Example

```go
package main

import (
  crypto "github.com/portalnesia/go-crypto"
  "fmt"
)

func main() {
  c := crypto.PortalnesiaCrypto("YOUR_SECRET_KEY")

  text := "Hello World"

  // Encrypted Text
  encrypted := c.Encrypt(text)
  fmt.Printf("Encrypted Text: %s",encrypted)

  // Decrypted Text
  decrypted := c.Decrypt(encrypted)
  fmt.Printf("Dncrypted Text: %s",decrypted)
}
```

## Go References
[pkg.go.dev/github.com/portalnesia/go-crypto](https://pkg.go.dev/github.com/portalnesia/go-crypto)