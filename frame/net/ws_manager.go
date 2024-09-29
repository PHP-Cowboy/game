package net

import (
	"common/logs"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WsManager struct {
	sync.RWMutex
	websocketUpgrade *websocket.Upgrader
}

func NewWsManager() *WsManager {
	return &WsManager{}
}

func (m *WsManager) Run(addr string) {

	http.HandleFunc("/", m.serveWS)
}

func (m *WsManager) serveWS(w http.ResponseWriter, r *http.Request) {
	if m.websocketUpgrade == nil {
		m.websocketUpgrade = &upgrade
	}

	wsConn, err := m.websocketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		logs.Error("websocketUpgrade.Upgrade err:%v", err)
		return
	}

	client := NewWsConnection(wsConn, m)

	client.Run()
}
