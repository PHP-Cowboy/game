package app

import (
	"center/router"
	"common/logs"
	"frame/center"
	"frame/core/resource"
	"log"
	"os"
	"os/signal"
)

func Run(serverId string) (err error) {
	logs.InitLog(serverId)

	c := center.NewCenter()

	c.RegisterHandler(router.Register(resource.NewManager()))

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	go func() {
		c.Run(serverId)
	}()

	<-quit
	log.Println("server is shutting down...")

	c.Close()

	close(done)
	log.Println("server stopped")

	return nil

}
