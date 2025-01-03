package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/augustus281/trackingcoin/internal/middlewares"
	"github.com/augustus281/trackingcoin/internal/wire"
)

type AssetRouterGroup struct {
	AssetRouter
}

type AssetRouter struct{}

func (r *AssetRouter) InitAssetRouter(route *gin.RouterGroup) {
	assetHandler, _ := wire.InitAssetRouterHandler()

	assetRouterPrivate := route.Group("/assets")
	{
		assetRouterPrivate.POST("/:asset_id/follow", middlewares.AuthenMiddleware(), assetHandler.FollowAsset)
		assetRouterPrivate.POST("/:asset_id/unfollow", middlewares.AuthenMiddleware(), assetHandler.UnfollowAsset)
	}
}
