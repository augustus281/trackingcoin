package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/augustus281/trackingcoin/internal/wire"
)

type CMCRouter struct{}

func (r *CMCRouter) InitCMCRouter(route *gin.RouterGroup) {
	cmcHandler, _ := wire.InitCMCRouterHandler()
	cmcRouterPublic := route.Group("/cmc")
	{
		cmcRouterPublic.GET("/market-pair", cmcHandler.GetMarketPairFromCMC)
		cmcRouterPublic.GET("/quote-lastest", cmcHandler.GetQuoteLastestFromCMC)
	}
}
