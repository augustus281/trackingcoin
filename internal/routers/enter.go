package routers

import "github.com/augustus281/trackingcoin/internal/routers/auth"

type RouterGroup struct {
	Auth auth.AuthRouter
}

var RouterGroupApp = new(RouterGroup)
