package initialize

import (
	"fmt"

	"github.com/augustus281/trackingcoin/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	InitRedis()
	InitKafka()
	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
