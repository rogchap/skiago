#include "include/core/SkCanvas.h"

#include "sk_canvas.h"

#include "sk_types_priv.h"

void sk_canvas_draw_rect(sk_canvas_t* ccanvas, const sk_rect_t* crect, const sk_paint_t* cpaint) {
    AsCanvas(ccanvas)->drawRect(*AsRect(crect), *AsPaint(cpaint));
}

void sk_canvas_draw_text_blob (sk_canvas_t* ccanvas, sk_textblob_t* text, float x, float y, const sk_paint_t* cpaint) {
    AsCanvas(ccanvas)->drawTextBlob(AsTextBlob(text), x, y, *AsPaint(cpaint));
}

void sk_canvas_draw_image(sk_canvas_t* ccanvas, const sk_image_t* cimage, float x, float y, const sk_sampling_options_t* sampling, const sk_paint_t* cpaint) {
    AsCanvas(ccanvas)->drawImage(AsImage(cimage), x, y);//, *AsSamplingOptions(sampling), AsPaint(cpaint));
}

void sk_canvas_draw_path(sk_canvas_t* ccanvas, const sk_path_t* cpath, const sk_paint_t* cpaint) {
    AsCanvas(ccanvas)->drawPath(*AsPath(cpath), *AsPaint(cpaint));
}

void sk_canvas_translate(sk_canvas_t* ccanvas, float dx, float dy) {
    AsCanvas(ccanvas)->translate(dx, dy);
}
