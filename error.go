package avmediacodec

import (
	"fmt"
)

type ErrAMedia int

func (e ErrAMedia) Error() string {
	return fmt.Sprintf("AMedia_error_%d", int(e))
}
