package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"time"
)

func (c *challenger) getHash(data []byte, time time.Time) []byte {
	hash := hmac.New(sha256.New, c.secretKey)
	hash.Write(data)

	b := make([]byte, 16)
	binary.LittleEndian.PutUint64(b, uint64(time.UnixNano()))
	hash.Write(b)
	return hash.Sum(nil)
}
