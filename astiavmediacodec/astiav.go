package astiavmediacodec

import (
	"reflect"

	"github.com/asticode/go-astiav"
	"github.com/xaionaro-go/avmediacodec/types"
	"github.com/xaionaro-go/unsafetools"
)
import "C"

func CFromAVCodecContext(codecCtx *astiav.CodecContext) *types.CVoid {
	orig := reflect.ValueOf(codecCtx)
	unsafePtr := unsafetools.FieldByNameInValue(orig, "c").Elem().UnsafePointer()
	return (*types.CVoid)(unsafePtr)
}
