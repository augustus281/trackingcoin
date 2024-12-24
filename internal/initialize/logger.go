package initialize

import (
	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
