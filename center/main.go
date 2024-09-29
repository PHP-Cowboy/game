package center

import (
	"center/app"
	"log"
	"os"
)

func main() {
	serverId := "001"

	//启动服务端
	err := app.Run(serverId)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
