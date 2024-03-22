#include "include/core/SkData.h"

#include "sk_data.h"

#include "sk_types_priv.h"

sk_data_t* sk_data_new_from_stream(sk_stream_t* stream, size_t length) {
    return ToData(SkData::MakeFromStream(AsStream(stream), length).release());
}
