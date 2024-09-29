package net

import "github.com/gorilla/websocket"

type WsConnection struct {
}

func NewWsConnection(conn *websocket.Conn, manager *WsManager) *WsConnection {
	return &WsConnection{}
}

func (c *WsConnection) Run() {

}
