#include "sk_types.h"

CPP_BEGIN_GUARD

sk_image_t* sk_image_new_from_encoded(const sk_data_t* cdata);

int sk_image_get_width(const sk_image_t* cimage);
int sk_image_get_height(const sk_image_t* cimage);

CPP_END_GUARD

