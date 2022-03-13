package model

type Solution struct {
	Challenge    []byte
	Nonce        int64
	LeadingZeros int
	Timestamp    int64
}
