#include "include/core/SkPath.h"

#include "sk_path.h"

#include "sk_types_priv.h"

sk_path_t* sk_path_new(void) {
    return ToPath(new SkPath());
}

void sk_path_move_to(sk_path_t* cpath, float x, float y) {
    AsPath(cpath)->moveTo(x, y);
}

void sk_path_line_to(sk_path_t* cpath, float x, float y) {
    AsPath(cpath)->lineTo(x, y);
}

void sk_path_arc_to(sk_path_t* cpath, const sk_rect_t* crect, float startAngle, float sweepAngle, bool forceMoveTo) {
    AsPath(cpath)->arcTo(*AsRect(crect), startAngle, sweepAngle, forceMoveTo);
}
