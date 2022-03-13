package model

import "time"

type ChallengeSolveData struct {
	AuthCode         []byte
	Data             []byte
	Time             time.Time
	Nonce            int64
	PrependZeroCount int64
}
