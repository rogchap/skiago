package skia

// #include "sk_data.h"
import "C"

import (
	"io"
	"syscall"
	"unsafe"
)

type Data struct {
	handle *C.sk_data_t
}

// NewData will read all data and copy the specified data into a new data ref.
// Note: although a copy is made, the bytes read will be GC'd; and ownership of
// the data will be with the data ref.
func NewData(r io.Reader) *Data {
	// TODO: if we had a logger we could output the error, or should we return errors?
	b, _ := io.ReadAll(r)
	return NewDataWithCopy(b)
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

func NewDataFromFile(path string) *Data {
	ptr, _ := syscall.BytePtrFromString(path)
	handle := C.sk_data_new_from_file((*C.char)(unsafe.Pointer(ptr)))
	if handle == nil {
		return nil
	}

	return &Data{handle}
}

func NewDataWithCopy(b []byte) *Data {
	handle := C.sk_data_new_with_copy(unsafe.Pointer(unsafe.SliceData(b)), C.ulong(len(b)))
	if handle == nil {
		return nil
	}

	return &Data{handle}
}
