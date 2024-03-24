#include "include/core/SkData.h"

#include "sk_data.h"

#include "sk_types_priv.h"

sk_data_t* sk_data_new_with_copy(const void* src, size_t length) {
    return ToData(SkData::MakeWithCopy(src, length).release());
}

sk_data_t* sk_data_new_from_file(const char* path) {
    return ToData(SkData::MakeFromFileName(path).release());
}

sk_data_t* sk_data_new_from_stream(sk_stream_t* stream, size_t length) {
    return ToData(SkData::MakeFromStream(AsStream(stream), length).release());
}
