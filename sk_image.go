package skia

// #include "sk_image.h"
import "C"
import "io"

type Image struct {
	handle *C.sk_image_t
}

func NewImage(r io.Reader) *Image {
	data := NewData(r)
	if data == nil {
		return nil
	}
	return NewImageFromEncoded(data)
}

func NewImageFromFile(path string) *Image {
	data := NewDataFromFile(path)
	if data == nil {
		return nil
	}
	return NewImageFromEncoded(data)
}

func NewImageFromEncoded(d *Data) *Image {
	return &Image{
		handle: C.sk_image_new_from_encoded(d.handle),
	}
}

func (i *Image) Width() int {
	return int(C.sk_image_get_width(i.handle))
}

func (i *Image) Height() int {
	return int(C.sk_image_get_height(i.handle))
}
