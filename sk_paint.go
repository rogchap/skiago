package skia

//#include "sk_paint.h"
import "C"

import "image/color"

type PaintStyle C.sk_paint_style_t

const (
	PaintStyleFill PaintStyle = iota
	PaintStyleStroke
	PaintStyleStrokeAndFill
)

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

func (p *Paint) SetStrokeWidth(w float32) {
	C.sk_paint_set_stroke_width(p.handle, C.float(w))
}

func (p *Paint) SetStyle(s PaintStyle) {
	C.sk_paint_set_style(p.handle, C.sk_paint_style_t(s))
}

func (p *Paint) SetAntiAlias(aa bool) {
	C.sk_paint_set_antialias(p.handle, C._Bool(aa))
}

func (p *Paint) SetStrokeCap(c StrokeCap) {
	C.sk_paint_set_stroke_cap(p.handle, c)
}
