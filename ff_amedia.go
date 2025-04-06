package avmediacodec

// #include <stdlib.h>
// #include <stdint.h>
// #include "3rdparty/ffmpeg/mediacodec_wrapper.h"
/*

typedef void *AVBSFContext;
typedef void *AVMutex;
typedef void *AVCond;
typedef void *AVFifo;

// copied from https://github.com/FFmpeg/FFmpeg/blob/master/libavcodec/mediacodecenc.c#L64-L105

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
*/
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avmediacodec/types"
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
) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.ff_AMediaFormat_setInt32((*C.FFAMediaFormat)(fmt), cname, C.int32_t(value))
	return nil
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
