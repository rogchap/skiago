#include "sk_types.h"

CPP_BEGIN_GUARD

sk_path_t* sk_path_new(void);

void sk_path_move_to(sk_path_t*, float x, float y);
void sk_path_line_to(sk_path_t*, float x, float y);
void sk_path_arc_to(sk_path_t* cpath, const sk_rect_t* crect, float startAngle, float sweepAngle, bool forceMoveTo);

CPP_END_GUARD

