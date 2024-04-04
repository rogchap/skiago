#ifndef sk_types_DEFINED
#define sk_types_DEFINED

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
    #define CPP_BEGIN_GUARD     extern "C" {
    #define CPP_END_GUARD       }
#else
    #include <stdbool.h>
    #define CPP_BEGIN_GUARD
    #define CPP_END_GUARD
#endif

CPP_BEGIN_GUARD

typedef uint32_t sk_color_t;

typedef struct sk_pixmap_t sk_pixmap_t;
typedef struct sk_wstream_filestream_t sk_wstream_filestream_t;
typedef struct sk_gostream_t sk_gostream_t;
typedef struct sk_gowstream_t sk_gowstream_t;
typedef struct sk_document_t sk_document_t;
typedef struct sk_stream_t sk_stream_t;
typedef struct sk_wstream_t sk_wstream_t;
typedef struct sk_canvas_t sk_canvas_t;
typedef struct sk_paint_t sk_paint_t;
typedef struct sk_textblob_t sk_textblob_t;
typedef struct sk_textblob_builder_t sk_textblob_builder_t;
typedef struct sk_font_t sk_font_t;
typedef struct sk_typeface_t sk_typeface_t;
typedef struct sk_fontmgr_t sk_fontmgr_t;
typedef struct sk_fontstyle_t sk_fontstyle_t;
typedef struct sk_data_t sk_data_t;
typedef struct sk_image_t sk_image_t;
typedef struct sk_string_t sk_string_t;
typedef struct sk_path_t sk_path_t;

typedef struct {
    float   left;
    float   top;
    float   right;
    float   bottom;
} sk_rect_t;

typedef struct {
    void* glyphs;
    void* pos;
    void* utf8text;
    void* clusters;
} sk_textblob_builder_runbuffer_t;

typedef struct {
    float   x;
    float   y;
} sk_point_t;

typedef enum {
    UTF8_SK_TEXT_ENCODING,
    UTF16_SK_TEXT_ENCODING,
    UTF32_SK_TEXT_ENCODING,
    GLYPH_ID_SK_TEXT_ENCODING
} sk_text_encoding_t;

typedef enum {
    UPRIGHT_SK_FONT_STYLE_SLANT = 0,
    ITALIC_SK_FONT_STYLE_SLANT  = 1,
    OBLIQUE_SK_FONT_STYLE_SLANT = 2,
} sk_font_style_slant_t;

typedef struct {
    int     fMaxAniso;
    bool    fUseCubic;
    // TODO: add missing fields
} sk_sampling_options_t;

typedef struct {
    int16_t  fTimeZoneMinutes;
    uint16_t fYear;
    uint8_t  fMonth;
    uint8_t  fDayOfWeek;
    uint8_t  fDay;
    uint8_t  fHour;
    uint8_t  fMinute;
    uint8_t  fSecond;
} sk_pdf_datetime_t;

typedef struct {
    sk_string_t*        fTitle;
    sk_string_t*        fAuthor;
    sk_string_t*        fSubject;
    sk_string_t*        fKeywords;
    sk_string_t*        fCreator;
    sk_string_t*        fProducer;
    sk_pdf_datetime_t*  fCreation;
    sk_pdf_datetime_t*  fModified;
    float               fRasterDPI;
    bool                fPDFA;
    int                 fEncodingQuality;
} sk_document_pdf_metadata_t;

typedef enum {
    FILL_SK_PAINT_STYLE,
    STROKE_SK_PAINT_STYLE,
    STROKE_AND_FILL_SK_PAINT_STYLE,
} sk_paint_style_t;

typedef enum {
    BUTT_SK_STROKE_CAP,
    ROUND_SK_STROKE_CAP,
    SQUARE_SK_STROKE_CAP
} sk_stroke_cap_t;

CPP_END_GUARD

#endif
