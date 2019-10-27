package domain

type ClientSource uint8

const (
	Manual ClientSource = 1 + iota
	Terminal
	Telegram
	WhatsUp
)
