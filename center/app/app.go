package app

import (
	"common/logs"
	"frame/center"
	"frame/core/resource"
)

func Run(serverId string) (err error) {
	logs.InitLog(serverId)

	c := center.NewCenter()

	m := resource.NewManager()

	return
}
