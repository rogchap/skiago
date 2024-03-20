package skia

//#include "sk_paint.h"
import "C"

import "image/color"

type Paint struct {
	handle *C.sk_paint_t
}

func NewPaint(c color.RGBA) *Paint {
	p := C.sk_paint_new()
	C.sk_paint_set_color(p, colorSetRGBA(c))
	return &Paint{p}
}

func colorSetRGBA(c color.RGBA) C.sk_color_t {
	return C.sk_color_t((uint32(c.A) << 24) | (uint32(c.R) << 16) | (uint32(c.G) << 8) | (uint32(c.B)))
}
