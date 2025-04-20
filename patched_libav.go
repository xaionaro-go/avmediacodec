//go:build patched_libav
// +build patched_libav

package avmediacodec

// #include "mediacodec_essentials.h"
import "C"

func (codec *FFAMediaCodec) SetParametersPatchedLibAV(avmFmt *FFAMediaFormat) error {
	status := C.ff_AMediaCodec_setParameters(codec.CPointer(), avmFmt.CPointer())
	if status != C.AMEDIA_OK {
		return ErrAMedia(status)
	}
	return nil
}
