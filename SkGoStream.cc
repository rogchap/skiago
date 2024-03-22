#include "SkGoStream.h"
#include "_cgo_export.h"

#include <iostream>

size_t SkGoStream::read(void* buffer, size_t size) {
    auto res = goReaderRead(fGoReader, buffer, size);
    fIsAtEnd = res.r1;
    // TODO: should we use the number of bytes read (res.r0) as a way to implement getPosition()?
    return size;
}

bool SkGoWStream::write(const void* buffer, size_t size) {
    // TODO: fBytesWritten should come from the return of the Go function
    fBytesWritten += size;

    goWriterWrite(fGoWriter, (void*)buffer, size);

    // TODO: return false if the Go code errors
    return true;
}
