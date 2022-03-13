package utils

import (
	"encoding/binary"
)

func Int64ToBytes(i int64) []byte {
	b := make([]byte, 16)
	binary.LittleEndian.PutUint64(b, uint64(i))

	return b
}
