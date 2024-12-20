package handler

import (
	"encoding/json"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
	"go-xianyu/internal/service"
	"go-xianyu/pkg/qiniu"
	"net/http"
	"strconv"

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
	var req v1.CreatePostRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	post := model.Post{
		Title:  req.Title,
		Info:   req.Info,
		Price:  req.Price,
		UserId: req.UserId,
		Img:    req.Img,
	}
	if err := h.postService.CreatePost(&post); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// GetPostListByPage godoc
//
//	@Summary	分页获取二手信息
//	@Schemes
//	@Description
//	@Tags		二手信息模块
//	@Accept		json
//	@Produce	json
//	@Param		pageNum	query	int		true	"page number"
//	@Param		pageSize	query	int		true	"page size"
//	@Security	Bearer
//	@Success	200		{object}	[]model.Post
//	@Router		/posts [get]
func (h *PostHandler) GetPostListByPage(ctx *gin.Context) {

	pageNum, err1 := strconv.Atoi(ctx.Query("pageNum"))
	pageSize, err2 := strconv.Atoi(ctx.Query("pageSize"))

	if err1 != nil || err2 != nil {
		v1.HandleError(ctx, http.StatusBadRequest, nil, "params invalid")
		return
	}

	data, err := h.postService.GetPostListByPage(pageNum, pageSize)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, nil, "service failed")
		return
	}

	v1.HandleSuccess(ctx, data)
}

// GetQiNiuToken godoc
//
//	@Summary	获取七牛云token
//	@Schemes
//	@Description
//	@Tags		七牛云模块
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Success	200		{object}	string
//	@Router		/qiniu/token [get]
func (h *PostHandler) GetQiNiuToken(ctx *gin.Context) {
	token := qiniu.GetToken()
	v1.HandleSuccess(ctx, token)
}

// GetPostInfo godoc
//
//	@Summary	获取二手信息
//	@Schemes
//	@Description
//	@Tags		二手信息模块
//	@Accept		json
//	@Produce	json
//	@Param		postId	query	int		true	"postId"
//	@Success	200		{object}	model.Post
//	@Router		/post/info [get]
func (h *PostHandler) GetPostInfo(ctx *gin.Context) {
	postId, _ := strconv.Atoi(ctx.Query("postId"))
	post := h.postService.GetPostInfo(uint(postId))
	v1.HandleSuccess(ctx, post)
}
