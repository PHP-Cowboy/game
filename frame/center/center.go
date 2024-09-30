package center

import "frame/net"

type Center struct {
	isRun     bool
	handlers  net.LogicHandler
	wsManager *net.WsManager
}

func NewCenter() *Center {
	return &Center{
		handlers: make(net.LogicHandler),
	}
}

func (c *Center) Run(serverId string) {
	if c.isRun {
		return
	}

}

func (c *Center) RegisterHandler(handlers net.LogicHandler) {
	c.handlers = handlers
}

func (c *Center) Close() {
	if !c.isRun {
		return
	}
	c.wsManager.Close()
}
