package handler

import (
	"encoding/json"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
	"go-xianyu/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	*Handler
	postService service.PostService
}

func NewPostHandler(
	handler *Handler,
	postService service.PostService,
) *PostHandler {
	return &PostHandler{
		Handler:     handler,
		postService: postService,
	}
}

func (h *PostHandler) GetPost(ctx *gin.Context) {

}

// CreatePost godoc
//
//	@Summary	发布二手信息
//	@Schemes
//	@Description
//	@Tags		二手信息模块
//	@Accept		json
//	@Produce	json
//	@Param		request	body		model.Post	true	"params"
//	@Security	Bearer
//	@Success	200
//	@Router		/post [post]
func (h *PostHandler) CreatePost(ctx *gin.Context) {
	var post model.Post

	if err := json.NewDecoder(ctx.Request.Body).Decode(&post); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err := h.postService.CreatePost(&post); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
