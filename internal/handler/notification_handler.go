package handler

import (
	"net/http"

	"github.com/augustus281/trackingcoin/internal/dto"
	"github.com/augustus281/trackingcoin/internal/service/notification"
	"github.com/gin-gonic/gin"
)

type NotifyHandler struct {
	service notification.INotifyService
}

func NewNotifyHandler(service notification.INotifyService) *NotifyHandler {
	return &NotifyHandler{
		service: service,
	}
}

func (h *NotifyHandler) SendNotification(ctx *gin.Context) {
	var notification dto.Notification
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Send(ctx, notification.Subject, notification.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func (h *NotifyHandler) ReceiveNotification(ctx *gin.Context) {
	err := h.service.Receive(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
