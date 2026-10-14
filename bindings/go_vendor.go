//go:build dummy_for_vendor

// This file is a workaround for `go mod vendor` so that
// the C and assembly source files in this directory
// are treated as part of a Go package and thus vendored.
package bindings

import (
	"embed"

	// Ensure that the build assembly is included
	_ "github.com/supranational/blst/build"
)

//go:embed ./*
var _ embed.FS
