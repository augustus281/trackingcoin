package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/augustus281/trackingcoin/internal/wire"
)

type AuthRouterGroup struct {
	AuthRouter
}

type AuthRouter struct{}

func (r *AuthRouter) InitAuthRouter(route *gin.RouterGroup) {
	authHandler, _ := wire.InitAuthRouterHandler()

	authRouterPublic := route.Group("/auth")
	{
		authRouterPublic.POST("/register", authHandler.Register)
		authRouterPublic.POST("/login", authHandler.Login)
	}
}
