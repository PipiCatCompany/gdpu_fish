package handler

import (
	"github.com/gin-gonic/gin"
	"go-xianyu/internal/service"
)

type MessageHandler struct {
	*Handler
	messageService service.MessageService
}

func NewMessageHandler(
    handler *Handler,
    messageService service.MessageService,
) *MessageHandler {
	return &MessageHandler{
		Handler:      handler,
		messageService: messageService,
	}
}

func (h *MessageHandler) GetMessage(ctx *gin.Context) {

}
