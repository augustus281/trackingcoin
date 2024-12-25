package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/augustus281/trackingcoin/internal/dto"
	repository "github.com/augustus281/trackingcoin/internal/repository/user"
)

type IAuthService interface {
	Register(ctx *gin.Context, email, password string) (int, error)
	Login(ctx *gin.Context, email, password string) (dto.LoginResponse, error)
}

type authService struct {
	repo repository.IUserRepository
}

func NewAuthService(repo repository.IUserRepository) IAuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(ctx *gin.Context, email, password string) (int, error) {
	return http.StatusOK, nil
}

func (s *authService) Login(ctx *gin.Context, email, password string) (dto.LoginResponse, error) {
	return dto.LoginResponse{}, nil
}
