package avmediacodec

/*
#include <stdlib.h>
#include <stdint.h>

typedef void AVClass;

// copied from https://github.com/FFmpeg/FFmpeg/blob/2471b22023756df145d24f9dcb00a56eaa121859/libavcodec/mediacodec_wrapper.h#L62-L88

typedef struct FFAMediaFormat FFAMediaFormat;

struct FFAMediaFormat {
    const AVClass *class;

    FFAMediaFormat *(*create)(void);
    int (*delete)(FFAMediaFormat *);

    char* (*toString)(FFAMediaFormat* format);

    int (*getInt32)(FFAMediaFormat* format, const char *name, int32_t *out);
    int (*getInt64)(FFAMediaFormat* format, const char *name, int64_t *out);
    int (*getFloat)(FFAMediaFormat* format, const char *name, float *out);
    int (*getBuffer)(FFAMediaFormat* format, const char *name, void** data, size_t *size);
    int (*getString)(FFAMediaFormat* format, const char *name, const char **out);
    // NDK only, introduced in API level 28
    int (*getRect)(FFAMediaFormat *, const char *name,
                   int32_t *left, int32_t *top, int32_t *right, int32_t *bottom);

    void (*setInt32)(FFAMediaFormat* format, const char* name, int32_t value);
    void (*setInt64)(FFAMediaFormat* format, const char* name, int64_t value);
    void (*setFloat)(FFAMediaFormat* format, const char* name, float value);
    void (*setString)(FFAMediaFormat* format, const char* name, const char* value);
    void (*setBuffer)(FFAMediaFormat* format, const char* name, void* data, size_t size);
    // NDK only, introduced in API level 28
    void (*setRect)(FFAMediaFormat*, const char* name,
                    int32_t left, int32_t top, int32_t right, int32_t bottom);
};

// copied from https://github.com/FFmpeg/FFmpeg/blob/2471b22023756df145d24f9dcb00a56eaa121859/libavcodec/mediacodec_wrapper.h#L135-L138

static inline void ff_AMediaFormat_setInt32(FFAMediaFormat* format, const char* name, int32_t value)
{
    format->setInt32(format, name, value);
}

*/
import "C"
import "unsafe"

type FFAMediaFormat C.FFAMediaFormat

func CWrapFFAMediaFormat(ptr *C.void) *C.FFAMediaFormat {
	return (*C.FFAMediaFormat)(unsafe.Pointer(ptr))
}

func WrapFFAMediaFormat(ptr *C.void) *FFAMediaFormat {
	return (*FFAMediaFormat)(CWrapFFAMediaFormat(ptr))
}

func (fmt *FFAMediaFormat) GetRawClass() unsafe.Pointer {
	return (*C.FFAMediaFormat)(fmt).class
}
func (fmt *FFAMediaFormat) GetRawCreate() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).create)
}
func (fmt *FFAMediaFormat) GetRawDelete() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).delete)
}
func (fmt *FFAMediaFormat) GetRawToString() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).toString)
}
func (fmt *FFAMediaFormat) GetRawGetInt32() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).getInt32)
}
func (fmt *FFAMediaFormat) GetRawGetInt64() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).getInt64)
}
func (fmt *FFAMediaFormat) GetRawGetFloat() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).getFloat)
}
func (fmt *FFAMediaFormat) GetRawGetBuffer() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).getBuffer)
}
func (fmt *FFAMediaFormat) GetRawGetString() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).getString)
}
func (fmt *FFAMediaFormat) GetRawGetRect() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).getRect)
}
func (fmt *FFAMediaFormat) GetRawSetInt32() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).setInt32)
}
func (fmt *FFAMediaFormat) GetRawSetInt64() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).setInt64)
}
func (fmt *FFAMediaFormat) GetRawSetFloat() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).setFloat)
}
func (fmt *FFAMediaFormat) GetRawSetString() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).setString)
}
func (fmt *FFAMediaFormat) GetRawSetBuffer() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).setBuffer)
}
func (fmt *FFAMediaFormat) GetRawSetRect() *C.void {
	return (*C.void)((*C.FFAMediaFormat)(fmt).setRect)
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
