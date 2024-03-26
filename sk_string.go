package skia

//#include "sk_string.h"
import "C"
import "unsafe"

type String struct {
	handle *C.sk_string_t
}

func NewString() *String {
	return &String{
		handle: C.sk_string_new_empty(),
	}
}

func (s *String) String() string {
	ptr := C.sk_string_get_c_str(s.handle)
	length := C.sk_string_get_size(s.handle)
	return unsafe.String((*byte)(unsafe.Pointer(ptr)), length)
}
