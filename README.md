# Base36

An implementation of base36 encoding and decoding for Go.

## Installing

Install the package by running `go get github.com/blakewilliams/go-base36`

## Usage

```go
import "github.com/blakewilliams/go-base36"

// Use the standard encoder
encoded := base36.StdEncoding.Encode(1234567890) // "kf12oi"
decoded, err := base36.StdEncoding.Decode(encoded) // 1234567890

// Create your own encoder
encoder := base36.NewEncoder("abcdefghijklmnopqrstuvwxyz0123456789")

encoded = encoder.Encode(1234567890) // "upbcys"
decoded, err = encoder.Decode("kf12oi") // 1234567890
```
