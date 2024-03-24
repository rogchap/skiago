#include "sk_types.h"

CPP_BEGIN_GUARD

sk_data_t* sk_data_new_with_copy(const void* src, size_t length);
sk_data_t* sk_data_new_from_file(const char* path);
sk_data_t* sk_data_new_from_stream(sk_stream_t* stream, size_t length);

CPP_END_GUARD

