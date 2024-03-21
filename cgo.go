package skia

// #cgo LDFLAGS: -lskia
// #cgo CXXFLAGS: -I${SRCDIR}/lib -std=c++17 -fvisibility-inlines-hidden -fno-exceptions -fno-rtti
// #cgo darwin LDFLAGS: -framework CoreFoundation -framework CoreGraphics -framework CoreText -framework CoreServices
// #cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/lib/darwin_arm64
import "C"

// These imports forces `go mod vendor` to pull in all the folders that
// contain Skia libraries and headers which otherwise would be ignored.
// DO NOT REMOVE
import (
	_ "rogchap.com/skia/lib/darwin_arm64"

	_ "rogchap.com/skia/lib/include/config"
	_ "rogchap.com/skia/lib/include/core"
	_ "rogchap.com/skia/lib/include/docs"
	_ "rogchap.com/skia/lib/include/ports"
	_ "rogchap.com/skia/lib/include/private"
	_ "rogchap.com/skia/lib/include/private/base"

	_ "rogchap.com/skia/lib/src/base"
)
