package initialize

import (
	"fmt"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/cronjob"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	//InitRedis()
	InitKafka()
	go cronjob.StartHourlyJob()
	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
