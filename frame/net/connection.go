package net

type Connection interface {
	Close()
	SendMessage([]byte) error
	GetSession() *Session
}

type MsgPack struct {
	Cid  string
	Body []byte
}
