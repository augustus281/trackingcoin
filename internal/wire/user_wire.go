//go:build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/augustus281/trackingcoin/internal/handler/auth"
	repository "github.com/augustus281/trackingcoin/internal/repository/user"
	service "github.com/augustus281/trackingcoin/internal/service/auth"
)

func InitAuthRouterHandler() (*auth.AuthHandler, error) {
	wire.Build(
		repository.NewUserRepo,
		service.NewAuthService,
		auth.NewAuthHandler,
	)
	return new(auth.AuthHandler), nil
}
