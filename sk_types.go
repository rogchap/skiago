package skia

//#include "sk_types.h"
import "C"

import "unsafe"

type Rect struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

func NewRect(x, y, w, h float32) Rect {
	return Rect{
		Left:   x,
		Top:    y,
		Right:  w,
		Bottom: h,
	}
}

func (r *Rect) cptr() *C.sk_rect_t {
	if r == nil {
		return nil
	}

	return (*C.sk_rect_t)(unsafe.Pointer(r))
}

type Point = C.sk_point_t

func NewPoint(x, y float32) Point {
	return C.sk_point_t{
		x: C.float(x),
		y: C.float(y),
	}
}

func (p Point) X() float32 {
	return float32(p.x)
}

func (p Point) Y() float32 {
	return float32(p.y)
}

type TextBlobBuilderRunBuffer = C.sk_textblob_builder_runbuffer_t

type FontStyleSlant = C.sk_font_style_slant_t

const (
	FontStyleSlantUright FontStyleSlant = iota
	FontStyleSlantItalic
	FontStyleSlantOblique
)

type FontStyleWeight int

const (
	FontStyleWeightHidden     FontStyleWeight = 0
	FontStyleWeightThin                       = 100
	FontStyleWeightExtraLight                 = 200
	FontStyleWeightLight                      = 300
	FontStyleWeightNormal                     = 400
	FontStyleWeightMedium                     = 500
	FontStyleWeightSemiBold                   = 600
	FontStyleWeightBold                       = 700
	FontStyleWeightExtraBold                  = 800
	FontStyleWeightBlack                      = 900
	FontStyleWeightExtraBlack                 = 1000
)

type FontStyleWidth int

const (
	FontStyleWidthUltraCondensed FontStyleWidth = 1
	FontStyleWidthExtraCondensed                = 2
	FontStyleWidthCondensed                     = 3
	FontStyleWidthSemiCondensed                 = 4
	FontStyleWidthNormal                        = 5
	FontStyleWidthSemiExpanded                  = 6
	FontStyleWidthExpanded                      = 7
	FontStyleWidthExtraExpanded                 = 8
	FontStyleWidthUltraExpanded                 = 9
)
