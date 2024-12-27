package routers

import (
	"github.com/augustus281/trackingcoin/internal/routers/asset"
	"github.com/augustus281/trackingcoin/internal/routers/auth"
	"github.com/augustus281/trackingcoin/internal/routers/coinmarketcap"
)

type RouterGroup struct {
	Auth  auth.AuthRouter
	CMC   coinmarketcap.CMCRouter
	Asset asset.AssetRouter
}

var RouterGroupApp = new(RouterGroup)
