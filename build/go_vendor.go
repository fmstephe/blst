//go:build dummy_for_vendor

// The !cgo has a critical function. This package contains no functionality and
// exists only to force the files in this directory to be preserved when this
// module is vendored. When vendored directories contain no go packages are
// stripped, which prevents these files from existing, which breaks the build.
// However, when the Go compiler compiles a package which contains a .c file
// but no cgo declarations using them it fails the compilation process. So we
// add the !cgo build directive to prevent the Go compiler from attempting to
// compile this package when cgo is true.
package build

// Ensure that the specific assembly directories are also included
import (
	"embed"

	_ "github.com/supranational/blst/build/cheri"
	_ "github.com/supranational/blst/build/coff"
	_ "github.com/supranational/blst/build/elf"
	_ "github.com/supranational/blst/build/mach-o"
	_ "github.com/supranational/blst/build/win64"
)

//go:embed ./*
var _ embed.FS
