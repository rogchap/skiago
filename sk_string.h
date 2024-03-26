#include "sk_types.h"

CPP_BEGIN_GUARD

sk_string_t* sk_string_new_empty(void);
sk_string_t* sk_string_new_with_copy(const char* src, size_t length);
size_t sk_string_get_size(const sk_string_t*);
const char* sk_string_get_c_str(const sk_string_t*);

CPP_END_GUARD

