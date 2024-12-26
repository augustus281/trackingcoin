//go:build wireinject

package wire

import (
	"github.com/google/wire"

	cmc "github.com/augustus281/trackingcoin/internal/handler/coinmarketcap"
	cmcRepo "github.com/augustus281/trackingcoin/internal/repository/coinmarketcap"
	cmcService "github.com/augustus281/trackingcoin/internal/service/coinmarketcap"
)

func InitCMCRouterHandler() (*cmc.CMCHandler, error) {
	wire.Build(
		cmcRepo.NewCMCRepo,
		cmcService.NewCMCService,
		cmc.NewCMCHandler,
	)
	return new(cmc.CMCHandler), nil
}
