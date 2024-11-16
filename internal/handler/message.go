package handler

import (
	"encoding/json"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/service"
	"net/http"
	"strconv"

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
//	@Summary	创建私聊 内置两个异步操作（Mysql和Cpp长连接）
//	@Schemes
//	@Description
//	@Tags		私聊模块
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

// GetMessageByPagination godoc
//
//	@Summary	分页获取私聊信息 （ godoc有问题）
//	@Schemes
//	@Description	根据聊天室ID、页码和页面大小分页获取私聊信息
//	@Tags		私聊模块
//	@Accept		json
//	@Produce	json
//	@Param		chatroomId  query   string	true	"聊天室ID"
//	@Param		pageNum		query	int		true	"页码"
//	@Param		pageSize	query	int		true	"页面大小"
//	@Security	Bearer
//	@Success	200		{object}   []v1.PostPaginationResponse
//	@Router		/msgs [get]
func (h *MessageHandler) GetMessageByPagination(ctx *gin.Context) {
	// chatroom
	chatroomId := ctx.Query("chatroomId")
	pageNum, err1 := strconv.Atoi(ctx.Query("pageNum"))
	pageSize, err2 := strconv.Atoi(ctx.Query("pageSize"))

	if err1 != nil || err2 != nil {
		v1.HandleError(ctx, http.StatusBadRequest, nil, "params invalid")
		return
	}

	data, err := h.messageService.GetMessageListByPage(pageSize, pageNum, chatroomId)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, nil, "service failed")
		return
	}

	v1.HandleSuccess(ctx, data)
}

// GetMessageChanelInfo godoc
//
//	@Summary	获取私聊频道信息
//	@Schemes
//	@Description	根据聊天室ID获取私聊频道信息
//	@Tags		私聊模块
//	@Accept		json
//	@Produce	json
//	@Param		chatroomId  query   string	true	"聊天室ID"
//	@Security	Bearer
//	@Success	200		{object} 	v1.MessageChanelResponse
//	@Router		/msgChanel [get]
func (h *MessageHandler) GetMessageChanelInfo(ctx *gin.Context) {
	// 获取聊天室信息
	chatroomId := ctx.Query("chatroomId")

	data, err := h.messageService.GetMessageChanelInfo(ctx, chatroomId)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, nil, "service failed")
		return
	}

	v1.HandleSuccess(ctx, data)
}
