#include "include/core/SkString.h"

#include "sk_string.h"

#include "sk_types_priv.h"

sk_string_t* sk_string_new_empty(void) {
    return ToString(new SkString());
}

sk_string_t* sk_string_new_with_copy(const char* src, size_t length) {
    return ToString(new SkString(src, length));
}

size_t sk_string_get_size(const sk_string_t* cstring) {
    return AsString(cstring)->size();
}

const char* sk_string_get_c_str(const sk_string_t* cstring) {
    return AsString(cstring)->c_str();
}
