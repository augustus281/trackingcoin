//go:build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/augustus281/trackingcoin/internal/handler"
	cmcRepo "github.com/augustus281/trackingcoin/internal/repository/coinmarketcap"
	cmcService "github.com/augustus281/trackingcoin/internal/service/coinmarketcap"
)

func InitCMCRouterHandler() (*handler.CMCHandler, error) {
	wire.Build(
		cmcRepo.NewCMCRepo,
		cmcService.NewCMCService,
		handler.NewCMCHandler,
	)
	return new(handler.CMCHandler), nil
}
