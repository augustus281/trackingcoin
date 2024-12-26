package util

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/augustus281/trackingcoin/global"
)

func HashPassword(password string) (string, error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		global.Logger.Error("generate from password failed ", zap.Error(err))
		return "", err
	}
	return string(hasedPassword), nil
}
