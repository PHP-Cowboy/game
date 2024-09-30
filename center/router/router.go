package router

import (
	"center/handler"
	"frame/core/resource"
	"frame/net"
)

func Register(m *resource.Manager) net.LogicHandler {
	h := make(net.LogicHandler)

	enterHandler := handler.NewEnterHandler()

	h["center.entry"] = enterHandler.Enter

	return h
}
