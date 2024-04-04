#include "include/core/SkPaint.h"

#include "sk_paint.h"

#include "sk_types_priv.h"

sk_paint_t* sk_paint_new(void) {
    return ToPaint(new SkPaint());
}

void sk_paint_set_color(sk_paint_t* cpaint, sk_color_t c) {
    AsPaint(cpaint)->setColor(c);
}

void sk_paint_set_stroke_width(sk_paint_t* cpaint, float width) {
    AsPaint(cpaint)->setStrokeWidth(width);
}

void sk_paint_set_style(sk_paint_t* cpaint, sk_paint_style_t style) {
    AsPaint(cpaint)->setStyle((SkPaint::Style)style);
}

void sk_paint_set_antialias(sk_paint_t* cpaint, bool aa) {
    AsPaint(cpaint)->setAntiAlias(aa);
}

void sk_paint_set_stroke_cap(sk_paint_t* cpaint, sk_stroke_cap_t ccap) {
    AsPaint(cpaint)->setStrokeCap((SkPaint::Cap)ccap);
}
