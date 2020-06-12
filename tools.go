// +build ignore

//go:generate go install github.com/go-bindata/go-bindata/v3/go-bindata
//go:generate go mod tidy

// Package installtools is a ghost library is intended only to help installing required tools locally.
// Run `go generate tools.go` from console before working with module.
package installtools

import (
	_ "github.com/go-bindata/go-bindata/v3"
)
