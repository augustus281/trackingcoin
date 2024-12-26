package routers

import (
	"github.com/augustus281/trackingcoin/internal/routers/auth"
	"github.com/augustus281/trackingcoin/internal/routers/coinmarketcap"
)

type RouterGroup struct {
	Auth auth.AuthRouter
	CMC  coinmarketcap.CMCRouter
}

var RouterGroupApp = new(RouterGroup)
