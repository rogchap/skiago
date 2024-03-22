#include <stdint.h>
#include "sk_types.h"

CPP_BEGIN_GUARD

sk_gostream_t* sk_gostream_t_new(uintptr_t r);
sk_gowstream_t* sk_gowstream_t_new(uintptr_t w);
sk_wstream_filestream_t* sk_filewstream_new(const char* path);


// testing
void sk_gostream_t_read(sk_gostream_t* s);

CPP_END_GUARD

