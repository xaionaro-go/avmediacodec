package avmediacodec

// #include "mediacodec_essentials.h"
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avcommon"
	"github.com/xaionaro-go/avcommon/types"
)

type FFAMediaFormat C.FFAMediaFormat

func CWrapFFAMediaFormat(ptr *types.CVoid) *C.FFAMediaFormat {
	return (*C.FFAMediaFormat)(unsafe.Pointer(ptr))
}

func WrapFFAMediaFormat(ptr *types.CVoid) *FFAMediaFormat {
	return (*FFAMediaFormat)(CWrapFFAMediaFormat(ptr))
}

func (fmt *FFAMediaFormat) CPointer() *C.FFAMediaFormat {
	return (*C.FFAMediaFormat)(fmt)
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

type FFAMediaCodec C.FFAMediaCodec

func CWrapFFAMediaCodec(ptr *types.CVoid) *C.FFAMediaCodec {
	return (*C.FFAMediaCodec)(unsafe.Pointer(ptr))
}

func WrapFFAMediaCodec(ptr *types.CVoid) *FFAMediaCodec {
	return (*FFAMediaCodec)(CWrapFFAMediaCodec(ptr))
}

func (codec *FFAMediaCodec) CPointer() *C.FFAMediaCodec {
	return (*C.FFAMediaCodec)(codec)
}

func (codec *FFAMediaCodec) Format() *FFAMediaFormat {
	cOutFormat := C.ff_AMediaCodec_getOutputFormat((*C.FFAMediaCodec)(codec))
	return (*FFAMediaFormat)(cOutFormat)
}

type AVCodecContext avcommon.AVCodecContext

func WrapAVCodecContext(ptr *types.CVoid) *AVCodecContext {
	return (*AVCodecContext)(avcommon.CWrapAVCodecContext(ptr))
}

func (avctx *AVCodecContext) PrivData() *MediaCodecEncContext {
	return WrapMediaCodecEncContext((*avcommon.AVCodecContext)(avctx).PrivData())
}

type MediaCodecEncContext C.MediaCodecEncContext

func CWrapMediaCodecEncContext(ptr *types.CVoid) *C.MediaCodecEncContext {
	return (*C.MediaCodecEncContext)(unsafe.Pointer(ptr))
}

func WrapMediaCodecEncContext(ptr *types.CVoid) *MediaCodecEncContext {
	return (*MediaCodecEncContext)(CWrapMediaCodecEncContext(ptr))
}

func (encCtx *MediaCodecEncContext) Codec() *FFAMediaCodec {
	return (*FFAMediaCodec)(encCtx.codec)
}
