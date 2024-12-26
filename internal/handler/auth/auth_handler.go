package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
	service "github.com/augustus281/trackingcoin/internal/service/auth"
)

type AuthHandler struct {
	service service.IAuthService
}

func NewAuthHandler(service service.IAuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var request dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "bind JSON request failed",
		})
		return
	}

	statusCode, err := h.service.Register(ctx, request.Email, request.Password)
	if err != nil {
		global.Logger.Error("register failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "register failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": "register successfully",
	})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "bind JSON request failed",
		})
		return
	}

	tokens, err := h.service.Login(ctx, request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
