package skia

//#include "sk_typeface.h"
import "C"

import (
	"syscall"
	"unsafe"
)

type Typeface struct {
	handle *C.sk_typeface_t
}

// font manager

type FontMgr struct {
	handle *C.sk_fontmgr_t
}

// TODO: Move function to arch spacific file
func NewSystemFontMgr() *FontMgr {
	return &FontMgr{
		handle: C.sk_fontmgr_create_mac_default(),
	}
}

func (f *FontMgr) CountFamiles() int {
	return int(C.sk_fontmgr_count_families(f.handle))
}

func (f *FontMgr) FamilyName(idx int) string {
	str := NewString()
	C.sk_fontmgr_get_family_name(f.handle, C.int(idx), str.handle)
	return str.String()
}

func (f *FontMgr) Families() []string {
	var fams []string

	for idx := 0; idx < f.CountFamiles(); idx++ {
		fams = append(fams, f.FamilyName(idx))
	}

	return fams
}

func (f *FontMgr) MatchFamily(familyName string) *Typeface {
	return f.MatchFamilyStyle(familyName, FontStyleNormal)
}

func (f *FontMgr) MatchFamilyStyle(familyName string, style *FontStyle) *Typeface {
	var sPtr *C.char = nil
	if familyName != "" {
		ptr, _ := syscall.BytePtrFromString(familyName)
		sPtr = (*C.char)(unsafe.Pointer(ptr))
	}

	match := C.sk_fontmgr_match_family_style(f.handle, sPtr, style.handle)
	if match == nil {
		return nil
	}

	return &Typeface{match}
}

func (f *FontMgr) CreateFromFile(path string, idx int) *Typeface {
	ptr, _ := syscall.BytePtrFromString(path)
	cptr := (*C.char)(unsafe.Pointer(ptr))
	tf := C.sk_fontmgr_create_from_file(f.handle, cptr, C.int(idx))
	if tf == nil {
		return nil
	}
	return &Typeface{tf}
}

// font style

type FontStyle struct {
	handle *C.sk_fontstyle_t
}

func NewFontStyle(weight FontStyleWeight, width FontStyleWidth, slant FontStyleSlant) *FontStyle {
	return &FontStyle{
		handle: C.sk_fontstyle_new(C.int(weight), C.int(width), slant),
	}
}

var FontStyleNormal = NewFontStyle(FontStyleWeightNormal, FontStyleWidthNormal, FontStyleSlantUright)
