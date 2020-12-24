package identcode3131

import (
	"encoding/base64"
	"encoding/binary"
)

// PackIdentCode concat given serialValue and randomKey into a int64 identifier code.
func PackIdentCode(serialValue, randomKey int32) (identCode int64) {
	identCode = int64(uint64(serialValue) | (uint64(randomKey) << 32))
	return
}

// UnpackIdentCode extract serialValue and randomKey from given identCode.
func UnpackIdentCode(identCode int64) (serialValue, randomKey int32) {
	c := uint64(identCode)
	serialValue = int32(c & 0x7FFFFFFF)
	randomKey = int32((c >> 32) & 0x7FFFFFFF)
	return
}

// PackIdentToken encode given serialValue and randomKey into identToken string.
// Binary value is packaged with base64.RawURLEncoding.
func PackIdentToken(serialValue, randomKey int32) (identToken string) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint32(buf[0:], uint32(randomKey))
	binary.LittleEndian.PutUint32(buf[4:], uint32(serialValue))
	identToken = base64.RawURLEncoding.EncodeToString(buf)
	return
}

// UnpackIdentToken extract serialValue and randomKey from given identToken.
func UnpackIdentToken(identToken string) (serialValue, randomKey int32, err error) {
	buf, err := base64.RawURLEncoding.DecodeString(identToken)
	if nil != err {
		return
	}
	if len(buf) != 8 {
		err = &ErrTokenSize{
			Size: len(buf),
		}
		return
	}
	randomKey = int32(binary.LittleEndian.Uint32(buf[0:]))
	serialValue = int32(binary.LittleEndian.Uint32(buf[4:]))
	return
}
