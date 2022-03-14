package model

type Challenge struct {
	Challenge    []byte `json:"challenge"`
	LeadingZeros int    `json:"leading_zeros"`
}
