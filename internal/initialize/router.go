package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	authRouter := routers.RouterGroupApp.Auth
	MainGroup := r.Group("/api/v1")
	{
		authRouter.InitAuthRouter(MainGroup)
	}
	return r
}
