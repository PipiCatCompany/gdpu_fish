package handler

import (
	"encoding/json"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
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
		Handler:        handler,
		messageService: messageService,
	}
}

func (h *MessageHandler) GetMessage(ctx *gin.Context) {

}

// CreateMessage godoc
//
//	@Summary	创建留言
//	@Schemes
//	@Description
//	@Tags		留言模块
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.CreateMessageRequest	true	"params"
//	@Security	Bearer
//	@Success	200
//	@Router		/message [post]
func (h *MessageHandler) CreateMessage(ctx *gin.Context) {
	var req v1.CreateMessageRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	err := h.messageService.CreateMessage(req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
	}

	v1.HandleSuccess(ctx, nil)
}
