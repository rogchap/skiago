#include "sk_types_priv.h"
#include "sk_stream.h"

sk_gostream_t* sk_gostream_t_new(uintptr_t r) {
    return ToGoStream(new SkGoStream(r));
}

sk_gowstream_t* sk_gowstream_t_new(uintptr_t w) {
    return ToGoWStream(new SkGoWStream(w));
}

sk_wstream_filestream_t* sk_filewstream_new(const char* path) {
    return ToFileWStream(new SkFILEWStream(path));
}

// Testing
void sk_gostream_t_read(sk_gostream_t* s) {
    char* buffer = new char[10];
    AsGoStream(s)->read(buffer, 10);
}

