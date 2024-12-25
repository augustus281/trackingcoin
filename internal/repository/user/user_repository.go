package repository

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	db "github.com/augustus281/trackingcoin/database/sqlc"
	"github.com/augustus281/trackingcoin/global"
)

type IUserRepository interface {
	Create(ctx *gin.Context, email, hashedPassword string) (db.User, error)
	GetByEmail(ctx *gin.Context, email string) (db.User, error)
}

type userRepo struct{}

func NewUserRepo() IUserRepository {
	return &userRepo{}
}

func (r *userRepo) GetByEmail(ctx *gin.Context, email string) (db.User, error) {
	user, err := global.Db.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		global.Logger.Error("not found user by email ",
			zap.String("email", email),
			zap.Error(err),
		)
		return db.User{}, err
	}
	return user, nil
}

func (r *userRepo) Create(ctx *gin.Context, email, hashedPassword string) (db.User, error) {
	return db.User{}, nil
}
