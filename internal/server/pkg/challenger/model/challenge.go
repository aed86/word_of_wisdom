package model

import "time"

type Challenge struct {
	Data             []byte    `json:"data"`
	PrependZeroCount int       `json:"prependZeroCount"`
	AuthCodeHash     []byte    `json:"authCodeHash"`
	Time             time.Time `json:"time"`
}
