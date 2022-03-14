package model

type Solution struct {
	Hash         []byte
	Challenge    []byte
	Nonce        int64
	LeadingZeros int
	Timestamp    int64
}
