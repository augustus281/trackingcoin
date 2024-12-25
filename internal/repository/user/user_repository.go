package repository

import (
	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	GetByEmail(ctx *gin.Context, email string)
}

type userRepo struct{}

func NewUserRepo() IUserRepository {
	return &userRepo{}
}

func (r *userRepo) GetByEmail(ctx *gin.Context, email string) {}
 