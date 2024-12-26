package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/augustus281/trackingcoin/global"
)

func GenerateAccessToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Duration(global.Config.Jwt.Expiration)).Unix()
	claims["id"] = userID

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(global.Config.Jwt.AccessToken))
}

func GenerateRefreshToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Duration(global.Config.Jwt.RefreshExpiration)).Unix()
	claims["id"] = userID

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshToken.SignedString([]byte(global.Config.Jwt.RefreshToken))
}
