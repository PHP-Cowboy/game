package handler

import "frame/net"

type EnterHandler struct {
}

func (h *EnterHandler) Enter(session *net.Session, msg []byte) (any, error) {
	return nil, nil
}

func NewEnterHandler() *EnterHandler {
	return &EnterHandler{}
}
