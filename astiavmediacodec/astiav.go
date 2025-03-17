package astiavmediacodec

import (
	"reflect"
	"unsafe"

	"github.com/asticode/go-astiav"
	"github.com/xaionaro-go/unsafetools"
)
import "C"

func CFromAVCodecContext(codecCtx *astiav.CodecContext) *C.void {
	orig := reflect.ValueOf(codecCtx)
	field := unsafetools.FieldByNameInValue(orig, "c")
	unsafePtr := field.Convert(reflect.TypeOf(unsafe.Pointer(nil))).Interface().(unsafe.Pointer)
	return (*C.void)(unsafePtr)
}
