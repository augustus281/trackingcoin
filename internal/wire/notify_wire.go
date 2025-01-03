//go:build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/augustus281/trackingcoin/internal/handler"
	"github.com/augustus281/trackingcoin/internal/service/notification"
)

func InitNotifyRouterHandler() (*handler.NotifyHandler, error) {
	wire.Build(
		notification.NewNotifyService,
		handler.NewNotifyHandler,
	)
	return new(handler.NotifyHandler), nil
}
