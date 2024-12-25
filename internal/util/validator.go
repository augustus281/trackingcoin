package util

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
)

func Validate(r dto.RegisterRequest) error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		global.Logger.Error("invalid register request: ",
			zap.String("email", r.Email),
			zap.String("password", r.Password),
			zap.Error(err),
		)
		return err
	}
	return nil
}
