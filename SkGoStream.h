#ifndef SkGoWStream_h
#define SkGoWStream_h

#include "include/core/SkStream.h"

class SkGoStream : public SkStream {
public:
    SkGoStream(uintptr_t r): fIsAtEnd(false), fGoReader(r) {}

    size_t read(void* buffer, size_t size) override;
    bool isAtEnd() const override { return fIsAtEnd; };

private:
    typedef SkStream INHERITED;
    uintptr_t fGoReader;
    bool fIsAtEnd;
};

class SkGoWStream : public SkWStream {
public:
    SkGoWStream(uintptr_t w) : fBytesWritten(0), fGoWriter(w) {}

    bool write(const void* buffer, size_t size) override;
    size_t bytesWritten() const override { return fBytesWritten; };

private:
    typedef SkWStream INHERITED;
    uintptr_t fGoWriter;
    size_t fBytesWritten;
};

#endif
