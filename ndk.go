//go:build with_android_ndk
// +build with_android_ndk

package avmediacodec

import (
	"unsafe"

	"github.com/xaionaro-go/ndk/media"
)

// #include "mediacodec_essentials.h"
import "C"

func (fmt *FFAMediaFormat) GetMediaFormatNDK() *media.MediaFormat {
	return (*media.MediaFormat)(unsafe.Pointer((*C.FFAMediaFormatNdk)(unsafe.Pointer(fmt)).impl))
}

func (codec *FFAMediaCodec) GetMediaCodecNDK() *media.MediaCodec {
	return (*media.MediaCodec)(unsafe.Pointer((*C.FFAMediaCodecNdk)(unsafe.Pointer(codec)).impl))
}

func (codec *FFAMediaCodec) SetParametersNDK(avmFmt *FFAMediaFormat) error {
	return codec.GetMediaCodecNDK().SetParameters(avmFmt.GetMediaFormatNDK())
}
