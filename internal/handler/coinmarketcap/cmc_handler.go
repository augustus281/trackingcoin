package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	service "github.com/augustus281/trackingcoin/internal/service/coinmarketcap"
)

type CMCHandler struct {
	service service.ICMCService
}

func NewCMCHandler(service service.ICMCService) *CMCHandler {
	return &CMCHandler{
		service: service,
	}
}

func (h *CMCHandler) GetDetailFromCMC(ctx *gin.Context) {
	slug := ctx.Query("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid slug",
			"code":  http.StatusBadRequest,
		})
		return
	}
	currency, err := h.service.GetDetailFromCMC(ctx, slug)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get detail currency from cmc successfully",
		"data":    currency,
	})
}

func (h *CMCHandler) GetMarketPairFromCMC(ctx *gin.Context) {
	slug := ctx.Query("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid slug",
			"code":  http.StatusBadRequest,
		})
		return
	}
	marketPair, err := h.service.GetMarketPairFromCMC(ctx, slug)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get market pair from cmc successfully",
		"data":    marketPair,
	})
}

func (h *CMCHandler) GetQuoteLastestFromCMC(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
			"code":  http.StatusBadRequest,
		})
		return
	}
	cmcID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "parse id failed",
			"code":  http.StatusBadRequest,
		})
	}
	quoteLastest, err := h.service.GetQuoteLastestFromCMC(ctx, int(cmcID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get quote lastest from cmc successfully",
		"data":    quoteLastest,
	})
}
