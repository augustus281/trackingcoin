package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/augustus281/trackingcoin/internal/wire"
)

type NotifyRouter struct{}

type NotifyRouterGroup struct {
	NotifyRouter
}

func (r *NotifyRouter) InitNotifyRouter(route *gin.RouterGroup) {
	notifyHandler, _ := wire.InitNotifyRouterHandler()
	notifyRouterPublic := route.Group("/notification")
	{
		notifyRouterPublic.POST("/", notifyHandler.SendNotification)
		notifyRouterPublic.GET("/", notifyHandler.ReceiveNotification)
	}
}
