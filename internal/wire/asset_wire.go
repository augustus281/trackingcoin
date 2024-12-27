//go:build wireinject

package wire

import (
	"github.com/google/wire"

	assetHandler "github.com/augustus281/trackingcoin/internal/handler/asset"
	assetRepo "github.com/augustus281/trackingcoin/internal/repository/asset"
	assetService "github.com/augustus281/trackingcoin/internal/service/asset"
)

func InitAssetRouterHandler() (*assetHandler.AssetHandler, error) {
	wire.Build(
		assetRepo.NewAssetRepo,
		assetService.NewAssetService,
		assetHandler.NewAssetHandler,
	)
	return new(assetHandler.AssetHandler), nil
}
