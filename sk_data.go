package skia

// #include "sk_data.h"
import "C"

import (
	"io"
	"unsafe"
)

type Data struct {
	handle *C.sk_data_t
}

// NewData will read all data and copy the specified data into a new data ref.
func NewData(r io.Reader) *Data {
	// TODO: if we had a logger we could output the error, or should we return errors?
	b, _ := io.ReadAll(r)
	handle := C.sk_data_new_with_copy(unsafe.Pointer(unsafe.SliceData(b)), C.ulong(len(b)))
	if handle == nil {
		return nil
	}

	return &Data{handle}
}

// NewDataSize attempts to read size bytes into a new data ref.
// If the read succeeds, return the data, else return nil.
// Either way the reader's cursor may have been changed as a result of calling read().
func NewDataSize(r io.Reader, size int) *Data {
	s := newGoStream(r)
	handle := C.sk_data_new_from_stream((*C.sk_stream_t)(unsafe.Pointer(s.handle)), C.ulong(size))
	if handle == nil {
		return nil
	}

	return &Data{handle}
}
