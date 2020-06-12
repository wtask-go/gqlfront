// +build ignore

//go:generate go get -u github.com/go-bindata/go-bindata/v3/go-bindata
//go:generate go mod tidy

package gqlfront

import (
	_ "github.com/go-bindata/go-bindata/v3"
)
