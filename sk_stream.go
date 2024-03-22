package skia

//#include "sk_stream.h"
import "C"

import (
	"errors"
	"io"
	"syscall"
	"unsafe"
)

type goStream struct {
	handle *C.sk_gostream_t
	reader io.Reader
}

func newGoStream(r io.Reader) *goStream {
	s := &goStream{
		reader: r,
	}
	ptr := uintptr(unsafe.Pointer(s))
	s.handle = C.sk_gostream_t_new(C.uintptr_t(ptr))

	return s
}

//export goReaderRead
func goReaderRead(rPtr uintptr, buf unsafe.Pointer, size C.int) (C.int, bool, bool) {
	// woraround to supress go vet warning
	// Ref: https://github.com/golang/go/issues/58625
	s := (*goStream)(*(*unsafe.Pointer)(unsafe.Pointer(&rPtr)))
	b := unsafe.Slice((*byte)(buf), size)
	n, err := s.reader.Read(b)
	isAtEnd := errors.Is(err, io.EOF)
	hasError := false
	if err != nil && !isAtEnd {
		hasError = true
		// TODO: if we have a logger we could log the error
	}
	if n < int(size) {
		isAtEnd = true
	}
	return C.int(n), isAtEnd, hasError
}

// testing only
func (s *goStream) readTest() {
	C.sk_gostream_t_read(s.handle)
}

type goWStream struct {
	handle *C.sk_gowstream_t
	writer io.Writer
}

func newGoWStream(w io.Writer) *goWStream {
	ws := &goWStream{
		writer: w,
	}
	ptr := uintptr(unsafe.Pointer(ws))
	ws.handle = C.sk_gowstream_t_new(C.uintptr_t(ptr))

	return ws
}

//export goWriterWrite
func goWriterWrite(w uintptr, buffer unsafe.Pointer, size C.int) {
	// woraround to supress go vet warning
	// Ref: https://github.com/golang/go/issues/58625
	ws := (*goWStream)(*(*unsafe.Pointer)(unsafe.Pointer(&w)))
	// TODO return bytes written and if error return false back to C
	ws.writer.Write(C.GoBytes(buffer, size))
}

type FileStream struct {
	handle *C.sk_wstream_filestream_t
}

func NewFileStream(path string) *FileStream {
	ptr, _ := syscall.BytePtrFromString(path)
	return &FileStream{
		handle: C.sk_filewstream_new((*C.char)(unsafe.Pointer(ptr))),
	}
}
