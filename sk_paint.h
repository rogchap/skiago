#include "sk_types.h"

CPP_BEGIN_GUARD

sk_paint_t* sk_paint_new(void);
void sk_paint_set_color(sk_paint_t* cpaint, sk_color_t c);
void sk_paint_set_stroke_width(sk_paint_t*, float width);
void sk_paint_set_style(sk_paint_t*, sk_paint_style_t);
void sk_paint_set_antialias(sk_paint_t*, bool);
void sk_paint_set_stroke_cap(sk_paint_t*, sk_stroke_cap_t);

CPP_END_GUARD

