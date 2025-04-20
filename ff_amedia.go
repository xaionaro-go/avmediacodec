package avmediacodec

/*
#include <stdlib.h>
#include <stdint.h>
#include "3rdparty/ffmpeg/mediacodec_wrapper.h"

typedef void *AVBSFContext;
typedef void *AVMutex;
typedef void *AVCond;
typedef void *AVFifo;
typedef void *AMediaCodec;
typedef void *AMediaFormat;
typedef void *ANativeWindow;
typedef void *AVFifo;
typedef void *media_status_t;
typedef void *bool;

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

*/
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avmediacodec/types"
	"github.com/xaionaro-go/ndk/media"
)

type FFAMediaFormat C.FFAMediaFormat

func CWrapFFAMediaFormat(ptr *types.CVoid) *C.FFAMediaFormat {
	return (*C.FFAMediaFormat)(unsafe.Pointer(ptr))
}

func WrapFFAMediaFormat(ptr *types.CVoid) *FFAMediaFormat {
	return (*FFAMediaFormat)(CWrapFFAMediaFormat(ptr))
}

func (fmt *FFAMediaFormat) GetRawClass() *C.AVClass {
	return (*C.FFAMediaFormat)(fmt).class
}
func (fmt *FFAMediaFormat) GetRawCreate() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).create)
}
func (fmt *FFAMediaFormat) GetRawDelete() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).delete)
}
func (fmt *FFAMediaFormat) GetRawToString() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).toString)
}
func (fmt *FFAMediaFormat) GetRawGetInt32() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).getInt32)
}
func (fmt *FFAMediaFormat) GetRawGetInt64() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).getInt64)
}
func (fmt *FFAMediaFormat) GetRawGetFloat() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).getFloat)
}
func (fmt *FFAMediaFormat) GetRawGetBuffer() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).getBuffer)
}
func (fmt *FFAMediaFormat) GetRawGetString() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).getString)
}
func (fmt *FFAMediaFormat) GetRawGetRect() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).getRect)
}
func (fmt *FFAMediaFormat) GetRawSetInt32() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).setInt32)
}
func (fmt *FFAMediaFormat) GetRawSetInt64() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).setInt64)
}
func (fmt *FFAMediaFormat) GetRawSetFloat() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).setFloat)
}
func (fmt *FFAMediaFormat) GetRawSetString() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).setString)
}
func (fmt *FFAMediaFormat) GetRawSetBuffer() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).setBuffer)
}
func (fmt *FFAMediaFormat) GetRawSetRect() *types.CVoid {
	return (*types.CVoid)((*C.FFAMediaFormat)(fmt).setRect)
}

func (fmt *FFAMediaFormat) SetInt32(
	name string,
	value int32,
) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.ff_AMediaFormat_setInt32((*C.FFAMediaFormat)(fmt), cname, C.int32_t(value))
}

func (fmt *FFAMediaFormat) GetInt32(
	name string,
) (int32, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var value int32
	C.ff_AMediaFormat_getInt32((*C.FFAMediaFormat)(fmt), cname, (*C.int32_t)(&value))
	return value, nil
}

func (fmt *FFAMediaFormat) SetInt64(
	name string,
	value int64,
) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.ff_AMediaFormat_setInt64((*C.FFAMediaFormat)(fmt), cname, C.int64_t(value))
}

func (fmt *FFAMediaFormat) GetInt64(
	name string,
) (int64, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var value int64
	C.ff_AMediaFormat_getInt64((*C.FFAMediaFormat)(fmt), cname, (*C.int64_t)(&value))
	return value, nil
}

func (fmt *FFAMediaFormat) GetMediaFormatNDK() *media.MediaFormat {
	return (*media.MediaFormat)(unsafe.Pointer((*C.FFAMediaFormatNdk)(unsafe.Pointer(fmt)).impl))
}

type FFAMediaCodec C.FFAMediaCodec

func CWrapFFAMediaCodec(ptr *types.CVoid) *C.FFAMediaCodec {
	return (*C.FFAMediaCodec)(unsafe.Pointer(ptr))
}

func WrapFFAMediaCodec(ptr *types.CVoid) *FFAMediaCodec {
	return (*FFAMediaCodec)(CWrapFFAMediaCodec(ptr))
}

func (codec *FFAMediaCodec) Format() *FFAMediaFormat {
	cOutFormat := C.ff_AMediaCodec_getOutputFormat((*C.FFAMediaCodec)(codec))
	return (*FFAMediaFormat)(cOutFormat)
}

func (codec *FFAMediaCodec) GetMediaCodecNDK() *media.MediaCodec {
	return (*media.MediaCodec)(unsafe.Pointer((*C.FFAMediaCodecNdk)(unsafe.Pointer(codec)).impl))
}

func (codec *FFAMediaCodec) SetParameters(avmFmt *FFAMediaFormat) error {
	return codec.GetMediaCodecNDK().SetParameters(avmFmt.GetMediaFormatNDK())
}

type AVCodecContext C.AVCodecContext

func CWrapAVCodecContext(ptr *types.CVoid) *C.AVCodecContext {
	return (*C.AVCodecContext)(unsafe.Pointer(ptr))
}

func WrapAVCodecContext(ptr *types.CVoid) *AVCodecContext {
	return (*AVCodecContext)(CWrapAVCodecContext(ptr))
}

func (avctx *AVCodecContext) PrivData() *MediaCodecEncContext {
	return (*MediaCodecEncContext)(avctx.priv_data)
}

type MediaCodecEncContext C.MediaCodecEncContext

func (encCtx *MediaCodecEncContext) Codec() *FFAMediaCodec {
	return (*FFAMediaCodec)(encCtx.codec)
}
