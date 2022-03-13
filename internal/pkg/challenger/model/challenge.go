package model

type Challenge struct {
	ChallengeData []byte `json:"challenge_data"`
	LeadingZeros  int    `json:"leading_zeros"`
	Timestamp     int64  `json:"timestamp"`
}
