package handler

import (
	"github.com/gin-gonic/gin"
	"go-xianyu/internal/service"
)

type ChatroomHandler struct {
	*Handler
	chatroomService service.ChatroomService
}

func NewChatroomHandler(
    handler *Handler,
    chatroomService service.ChatroomService,
) *ChatroomHandler {
	return &ChatroomHandler{
		Handler:      handler,
		chatroomService: chatroomService,
	}
}

func (h *ChatroomHandler) GetChatroom(ctx *gin.Context) {

}
