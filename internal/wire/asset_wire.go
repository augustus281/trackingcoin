//go:build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/augustus281/trackingcoin/internal/handler"
	assetRepo "github.com/augustus281/trackingcoin/internal/repository/asset"
	assetService "github.com/augustus281/trackingcoin/internal/service/asset"
)

func InitAssetRouterHandler() (*handler.AssetHandler, error) {
	wire.Build(
		assetRepo.NewAssetRepo,
		assetService.NewAssetService,
		handler.NewAssetHandler,
	)
	return new(handler.AssetHandler), nil
}
