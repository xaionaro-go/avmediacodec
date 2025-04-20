#include <stdlib.h>
#include <stdint.h>
#include "3rdparty/ffmpeg/mediacodec_wrapper.h"
#include <media/NdkMediaCodec.h>

typedef void *AVBSFContext;
typedef void *AVMutex;
typedef void *AVCond;
typedef void *AVFifo;

// copied from https://github.com/FFmpeg/FFmpeg/blob/853e66a0726b0a9d6d6269a22f6f9b5be7763738/libavcodec/mediacodecenc.c#L64-L105

typedef struct MediaCodecEncContext {
    AVClass *avclass;
    FFAMediaCodec *codec;
    int use_ndk_codec;
    const char *name;
    FFANativeWindow *window;

    int fps;
    int width;
    int height;

    uint8_t *extradata;
    int extradata_size;
    int eof_sent;

    AVFrame *frame;
    AVBSFContext *bsf;

    int bitrate_mode;
    int level;
    int pts_as_dts;
    int extract_extradata;
    // Ref. MediaFormat KEY_OPERATING_RATE
    int operating_rate;
    int async_mode;

    AVMutex input_mutex;
    AVCond input_cond;
    AVFifo *input_index;

    AVMutex output_mutex;
    AVCond output_cond;
    int encode_status;
    AVFifo *async_output;

    int qp_i_min;
    int qp_p_min;
    int qp_b_min;
    int qp_i_max;
    int qp_p_max;
    int qp_b_max;
} MediaCodecEncContext;

// copied from https://github.com/FFmpeg/FFmpeg/blob/853e66a0726b0a9d6d6269a22f6f9b5be7763738/libavcodec/mediacodec_wrapper.c#L1835-L1866

typedef struct FFAMediaFormatNdk {
    FFAMediaFormat api;

    void *libmedia;
    AMediaFormat *impl;

    bool (*getRect)(AMediaFormat *, const char *name,
                    int32_t *left, int32_t *top, int32_t *right, int32_t *bottom);
    void (*setRect)(AMediaFormat *, const char *name,
                    int32_t left, int32_t top, int32_t right, int32_t bottom);
} FFAMediaFormatNdk;

typedef struct FFAMediaCodecNdk {
    FFAMediaCodec api;

    void *libmedia;
    AMediaCodec *impl;
    ANativeWindow *window;

    FFAMediaCodecOnAsyncNotifyCallback async_cb;
    void *async_userdata;

    // Available since API level 28.
    media_status_t (*getName)(AMediaCodec*, char** out_name);
    void (*releaseName)(AMediaCodec*, char* name);

    // Available since API level 26.
    media_status_t (*setInputSurface)(AMediaCodec*, ANativeWindow *);
    media_status_t (*signalEndOfInputStream)(AMediaCodec *);
    media_status_t (*setAsyncNotifyCallback)(AMediaCodec *,
            struct AMediaCodecOnAsyncNotifyCallback callback, void *userdata);
} FFAMediaCodecNdk;
