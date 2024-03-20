package skia

// #cgo LDFLAGS: -lskia
// #cgo CXXFLAGS: -I${SRCDIR}/lib -std=c++17 -fvisibility-inlines-hidden -fno-exceptions -fno-rtti
// #cgo darwin LDFLAGS: -framework CoreFoundation -framework CoreGraphics -framework CoreText -framework CoreServices
// #cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/lib/darwin_arm64
import "C"
