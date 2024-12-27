package extractor

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
)

type Extractor interface {
	GetUserID(ctx *gin.Context, token string) (int, error)
}

type extractor struct{}

func New() Extractor {
	return &extractor{}
}

func (t *extractor) GetUserID(ctx *gin.Context, token string) (int, error) {
	claim, err := t.ParseToken(token)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(claim.UserID)
}

func (t *extractor) ParseToken(token string) (*Token, error) {
	accessToken, _, err := new(jwt.Parser).ParseUnverified(token, &Token{})
	if err != nil {
		global.Logger.Error("error parse token ", zap.Error(err))
		return nil, err
	}
	claims, ok := accessToken.Claims.(*Token)
	if !ok {
		global.Logger.Error("extract claims error", zap.Error(err))
		return nil, err
	}
	return claims, nil
}
