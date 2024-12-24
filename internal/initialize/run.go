package initialize

import (
	"fmt"

	"github.com/augustus281/trackingcoin/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	if global.Config.Server.Mode != "release" {
		fmt.Println(serverAddr)
	}
	r.Run(serverAddr)
}
