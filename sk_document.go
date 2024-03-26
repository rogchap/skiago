package skia

/*
 #include <stdlib.h>
 #include "sk_document.h"

sk_pdf_datetime_t* allocate_pdf_datetime() {
    return (sk_pdf_datetime_t*)malloc(sizeof(sk_pdf_datetime_t));
}
*/
import "C"

import (
	"io"
	"time"
	"unsafe"
)

type Document struct {
	handle *C.sk_document_t

	w *goWStream
}

func NewDocument(w io.Writer) *Document {
	gws := newGoWStream(w)
	return &Document{
		handle: C.sk_document_create_pdf_from_stream((*C.sk_wstream_t)(unsafe.Pointer(gws.handle))),
		w:      gws,
	}
}

func NewDocumentWithMetadata(w io.Writer, metadata PDFMetadata) *Document {
	gws := newGoWStream(w)
	cMetadata := metadata.cMetadata()
	defer func() {
		C.free(unsafe.Pointer(cMetadata.fCreation))
		C.free(unsafe.Pointer(cMetadata.fModified))
	}()

	return &Document{
		handle: C.sk_document_create_pdf_from_stream_with_metadata((*C.sk_wstream_t)(unsafe.Pointer(gws.handle)), cMetadata),
		w:      gws,
	}
}

func (d *Document) BeginPage(width, height float32, content *Rect) *Canvas {
	if width == 0 || height == 0 {
		panic("skia: can't begin a page with zero width or height")
	}

	return &Canvas{
		handle: C.sk_document_begin_page(d.handle, C.float(width), C.float(height), content.cptr()),
	}
}

func (d *Document) EndPage() {
	C.sk_document_end_page(d.handle)
}

func (d *Document) Close() {
	C.sk_document_close(d.handle)
}

func (d *Document) Dispose() {
	C.sk_document_unref(d.handle)
	d.handle = nil
	d.w = nil
	d = nil
}

// Optional metadata to be passed into the NewDocument function.
type PDFMetadata struct {
	// The document's title.
	Title string
	// The name of the person who created the document.
	Author string
	// The subject of the document.
	Subject string
	// Keywords associated with the document.
	Keywords string
	// If the document was converted to PDF from another format, the name of the conforming product that created the
	// original document from which it was converted.
	Creator string
	// The product that is converting this document to PDF.
	Producer string
	// The date and time the document was created.
	Creation time.Time
	// The date and time the document was most recently modified.
	Modified time.Time

	// If true, include XMP metadata, a document UUID, and sRGB output intent information.
	// This adds length to the document and makes it non-reproducable, but are necessary features for PDF/A-2b conformance
	PDFA bool
	// The DPI (pixels-per-inch) at which features without native PDF support will be rasterized (e.g. draw image with
	// perspective, draw text with perspective, ...) A larger DPI would create a PDF that reflects the original intent
	// with better fidelity, but it can make for larger PDF files too, which would use more memory while rendering, and
	// it would be slower to be processed or sent online or to printer.
	RasterDPI float32
	// Encoding quality controls the trade-off between size and quality.
	// By default this is set to 101 percent, which corresponds to lossless encoding. If this value is set to a
	// value <= 100, and the image is opaque, it will be encoded (using JPEG) with that quality setting.
	EncodingQuality int

	// TODO: Add Compression Level
}

const DefaultRasterDPI = 72.0

func (m PDFMetadata) cMetadata() *C.sk_document_pdf_metadata_t {
	var cPDFMetadata C.sk_document_pdf_metadata_t

	if m.Title != "" {
		title := NewStringWithCopy(m.Title)
		cPDFMetadata.fTitle = title.handle
	}

	if m.Author != "" {
		author := NewStringWithCopy(m.Author)
		cPDFMetadata.fAuthor = author.handle
	}

	if m.Subject != "" {
		subject := NewStringWithCopy(m.Subject)
		cPDFMetadata.fSubject = subject.handle
	}

	if m.Keywords != "" {
		keywords := NewStringWithCopy(m.Keywords)
		cPDFMetadata.fKeywords = keywords.handle
	}

	if m.Creator != "" {
		creator := NewStringWithCopy(m.Creator)
		cPDFMetadata.fCreator = creator.handle
	}

	if m.Producer != "" {
		producer := NewStringWithCopy(m.Producer)
		cPDFMetadata.fProducer = producer.handle
	}

	if !m.Creation.IsZero() {
		creation := cDateTime(m.Creation)
		cPDFMetadata.fCreation = (*C.sk_pdf_datetime_t)(unsafe.Pointer(creation))
	}

	if !m.Modified.IsZero() {
		modified := cDateTime(m.Modified)
		cPDFMetadata.fModified = (*C.sk_pdf_datetime_t)(unsafe.Pointer(&modified))
	}

	cPDFMetadata.fPDFA = C._Bool(m.PDFA)

	if m.RasterDPI <= 0 {
		m.RasterDPI = DefaultRasterDPI
	}
	cPDFMetadata.fRasterDPI = C.float(m.RasterDPI)

	if m.EncodingQuality <= 0 {
		m.EncodingQuality = 101
	}
	cPDFMetadata.fEncodingQuality = C.int(m.EncodingQuality)

	return &cPDFMetadata
}

func cDateTime(t time.Time) *C.sk_pdf_datetime_t {
	cPDFDateTime := C.allocate_pdf_datetime()

	_, offset := t.Zone()
	cPDFDateTime.fTimeZoneMinutes = C.short(offset / 60)
	cPDFDateTime.fYear = C.ushort(t.Year())
	cPDFDateTime.fMonth = C.uchar(t.Month())
	cPDFDateTime.fDayOfWeek = C.uchar(t.Weekday())
	cPDFDateTime.fDay = C.uchar(t.Day())
	cPDFDateTime.fHour = C.uchar(t.Hour())
	cPDFDateTime.fMinute = C.uchar(t.Minute())
	cPDFDateTime.fSecond = C.uchar(t.Second())

	return cPDFDateTime
}
