package main

import (
	"center/app"
	"common/config"
	"log"
	"os"
)

func main() {
	serverId := "001"

	config.InitConfig("config.json")

	//启动服务端
	err := app.Run(serverId)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
