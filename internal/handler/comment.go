package handler

import (
	"go-xianyu/internal/service"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	*Handler
	commentService service.CommentService
}

func NewCommentHandler(
	handler *Handler,
	commentService service.CommentService,
) *CommentHandler {
	return &CommentHandler{
		Handler:        handler,
		commentService: commentService,
	}
}

func (h *CommentHandler) GetComment(ctx *gin.Context) {

}

// TODO
// f'发送评论' -> 需要穿openid 且认证才可以
// f'获取pid下的所有评论' -> 不需要openid
