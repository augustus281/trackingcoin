package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
	repository "github.com/augustus281/trackingcoin/internal/repository/user"
	"github.com/augustus281/trackingcoin/internal/util"
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
	user, err := s.repo.GetByEmail(ctx, email)
	if err == nil || user.Email == email {
		global.Logger.Error("Email is already register",
			zap.String("email", email),
			zap.Error(err),
		)
		return http.StatusInternalServerError, errors.New("email is already register")
	}

	_, err = s.repo.Create(ctx, email, password)
	if err != nil {
		global.Logger.Error("create user failed ",
			zap.String("email", email),
			zap.Error(err),
		)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *authService) Login(ctx *gin.Context, email, password string) (dto.LoginResponse, error) {
	var response = dto.LoginResponse{}
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		global.Logger.Error("user not exist",
			zap.Error(err),
			zap.String("email", email),
		)
		return response, err
	}

	if err := util.CheckPassword(password, user.HashedPassword); err != nil {
		global.Logger.Error("wrong password ", zap.Error(err))
		return response, err
	}

	accessToken, err := util.GenerateAccessToken(int(user.ID))
	if err != nil {
		global.Logger.Error("generate access token failed ", zap.Error(err))
		return response, err
	}
	response.AccessToken = accessToken

	refreshToken, err := util.GenerateRefreshToken(int(user.ID))
	if err != nil {
		global.Logger.Error("generate refresh token failed ", zap.Error(err))
		return response, err
	}
	response.RefreshToken = refreshToken

	return response, nil
}
