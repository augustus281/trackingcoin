package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
	service "github.com/augustus281/trackingcoin/internal/service/asset"
	"github.com/augustus281/trackingcoin/pkg/extractor"
)

type AssetHandler struct {
	service   service.IAssetService
	extractor extractor.Extractor
}

func NewAssetHandler(service service.IAssetService) *AssetHandler {
	return &AssetHandler{
		service:   service,
		extractor: extractor.New(),
	}
}

func (h *AssetHandler) FollowAsset(ctx *gin.Context) {
	userID, err := h.extractor.GetUserID(ctx, ctx.GetHeader("Authorization"))
	if err != nil {
		global.Logger.Error("failed to get userID", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get userID from token",
		})
		return
	}

	asset_id := ctx.Param("asset_id")
	if asset_id == "" {
		global.Logger.Error("asset_id is invalid")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid asset id",
		})
		return
	}
	assetID, err := strconv.ParseInt(asset_id, 10, 64)
	if err != nil {
		global.Logger.Error("parse asset_id to int failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	data, err := h.service.FollowAsset(ctx, userID, int(assetID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "follow asset failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "follow asset successfully",
		"data":    data,
	})
}

func (h *AssetHandler) UnfollowAsset(ctx *gin.Context) {
	userID, err := h.extractor.GetUserID(ctx, ctx.GetHeader("Authorization"))
	if err != nil {
		global.Logger.Error("failed to get userID", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get userID from token",
		})
		return
	}

	asset_id := ctx.Param("asset_id")
	if asset_id == "" {
		global.Logger.Error("asset_id is invalid")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid asset id",
		})
		return
	}
	assetID, err := strconv.ParseInt(asset_id, 10, 64)
	if err != nil {
		global.Logger.Error("parse asset_id to int failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.service.UnfollowAsset(ctx, userID, int(assetID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unfollow asset failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "unfollow asset successfully",
	})
}
