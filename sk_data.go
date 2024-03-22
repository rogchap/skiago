package skia

// #include "sk_data.h"
import "C"

import (
	"io"
	"unsafe"
)

const defaultReadByteSize = 1024

type Data struct {
	handle *C.sk_data_t
}

func NewData(r io.Reader) *Data {
	s := newGoStream(r)
	return &Data{
		handle: C.sk_data_new_from_stream((*C.sk_stream_t)(unsafe.Pointer(s.handle)), C.ulong(defaultReadByteSize)),
	}
}
