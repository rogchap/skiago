#include "SkGoStream.h"
#include "_cgo_export.h"

size_t SkGoStream::read(void* buffer, size_t size) {
    auto res = goReaderRead(fGoReader, buffer, size);
    fIsAtEnd = res.r1;
    return res.r0;
}

bool SkGoWStream::write(const void* buffer, size_t size) {
    // TODO: fBytesWritten should come from the return of the Go function
    fBytesWritten += size;

    goWriterWrite(fGoWriter, (void*)buffer, size);

    // TODO: return false if the Go code errors
    return true;
}
