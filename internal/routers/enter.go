package routers

type RouterGroup struct {
	Auth         AuthRouter
	CMC          CMCRouter
	Asset        AssetRouter
	Notification NotifyRouter
}

var RouterGroupApp = new(RouterGroup)
