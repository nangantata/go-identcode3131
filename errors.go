package identcode3131

import (
	"strconv"
)

// ErrTokenSize indicate given token size is not correct.
type ErrTokenSize struct {
	Size int
}

func (e ErrTokenSize) Error() string {
	return "[ErrTokenSize (expect 8, got " + strconv.FormatInt(int64(e.Size), 10) + ")]"
}
