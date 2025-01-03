//go:build wireinject

package wire

import (
	"github.com/google/wire"

	handler "github.com/augustus281/trackingcoin/internal/handler"
	repository "github.com/augustus281/trackingcoin/internal/repository/user"
	service "github.com/augustus281/trackingcoin/internal/service/auth"
)

func InitAuthRouterHandler() (*handler.AuthHandler, error) {
	wire.Build(
		repository.NewUserRepo,
		service.NewAuthService,
		handler.NewAuthHandler,
	)
	return new(handler.AuthHandler), nil
}
