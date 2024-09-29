package net

import "frame/protocol"

type HandlerFunc func(session Session, msg []byte) (any, error)

type LogicHandler map[string]HandlerFunc

type EventHandler func(packageType *protocol.Packet, c Connection) error
