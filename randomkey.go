package identcode3131

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// GenerateRandomKey create 31 bits random key.
// Boolean isSecure will be false when failed to generate enough bits from
// crypto/rand.
func GenerateRandomKey() (randomKey int32, isSecure bool) {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); nil != err {
		randomKey = int32(uint64(time.Now().UnixNano()) & 0x7FFFFFFF)
		return
	}
	randomKey = int32(binary.LittleEndian.Uint32(buf) & 0x7FFFFFFF)
	isSecure = true
	return
}

// GenerateNonZeroRandomKey creates non-zero 31 bits random key.
// Boolean isSecure will be false when failed to generate enough bits from
// crypto/rand or result in 0 after attempts.
func GenerateNonZeroRandomKey(maxAttempts int) (randomKey int32, isSecure bool) {
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if randomKey, isSecure = GenerateRandomKey(); (randomKey != 0) && isSecure {
			return
		}
	}
	randomKey, isSecure = GenerateRandomKey()
	if randomKey != 0 {
		return
	}
	return 1, false
}
