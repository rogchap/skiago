package skia

// #include "sk_path.h"
import "C"

type Path struct {
	handle *C.sk_path_t
}

func NewPath() *Path {
	return &Path{
		handle: C.sk_path_new(),
	}
}

func (p *Path) MoveTo(x, y float32) {
	C.sk_path_move_to(p.handle, C.float(x), C.float(y))
}

func (p *Path) LineTo(x, y float32) {
	C.sk_path_line_to(p.handle, C.float(x), C.float(y))
}

func (p *Path) ArcTo(rect Rect, startAngle, sweepAngle float32, forceMoveTo bool) {
	C.sk_path_arc_to(p.handle, rect.cptr(), C.float(startAngle), C.float(sweepAngle), C._Bool(forceMoveTo))
}
