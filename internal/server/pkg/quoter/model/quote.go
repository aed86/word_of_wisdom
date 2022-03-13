package model

type Quote struct {
	Author string `json:"author" csv:"Author"`
	Text   string `json:"text" csv:"Quote"`
}
