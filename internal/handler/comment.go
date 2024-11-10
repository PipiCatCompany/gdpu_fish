package handler

import (
	"encoding/json"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/service"
	"net/http"
	"strconv"

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

// CreateComment godoc
//
//	@Summary	创建评论
//	@Schemes
//	@Description
//	@Tags		评论模块
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.CreateCommentRequest	true	"params"
//	@Security	Bearer
//	@Success	200
//	@Router		/comment [post]
func (h *CommentHandler) CreateComment(ctx *gin.Context) {
	var req v1.CreateCommentRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err := h.commentService.CreateComment(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// GetCommentList godoc
//
//	@Summary	获取商品下所有评论
//	@Schemes
//	@Description
//	@Tags		评论模块
//	@Accept		json
//	@Produce	json
//	@Param		postId	query string	true	"商品postId"
//	@Success	200		{object} 	[]v1.CommentResponse
//	@Router		/comments [get]
func (h *CommentHandler) GetCommentList(ctx *gin.Context) {
	postIdStr := ctx.Query("postId")
	postId64, err := strconv.ParseUint(postIdStr, 10, 64)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	postId := uint(postId64)

	// 获取postId下所有评论
	commentList, err := h.commentService.GetCommentList(postId)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, commentList)
}
