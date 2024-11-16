package handler

import (
	"encoding/json"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/service"
	"go-xianyu/pkg/wx"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

// Register godoc
//
//	@Summary	用户注册
//	@Schemes
//	@Description	目前只支持邮箱登录
//	@Tags			用户模块
//	@Accept			json
//	@Produce		json
//	@Param			request	body		v1.RegisterRequest	true	"params"
//	@Success		200		{object}	v1.Response
//	@Router			/register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	req := new(v1.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.Register(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// Login godoc
//
//	@Summary	账号登录
//	@Schemes
//	@Description
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.LoginRequest	true	"params"
//	@Success	200		{object}	v1.LoginResponse
//	@Router		/login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var req v1.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	v1.HandleSuccess(ctx, v1.LoginResponseData{
		AccessToken: token,
	})
}

// GetProfile godoc
//
//	@Summary	获取用户信息
//	@Schemes
//	@Description
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Success	200	{object}	v1.GetProfileResponse
//	@Router		/user [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	v1.HandleSuccess(ctx, user)
}

// UpdateProfile godoc
//
//	@Summary	修改用户信息
//	@Schemes
//	@Description
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		request	body		v1.UpdateProfileRequest	true	"params"
//	@Success	200		{object}	v1.Response
//	@Router		/user [put]
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req v1.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// CreateUserBasic godoc
//
//	@Summary	使用openid创建用户
//	@Schemes
//	@Description
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.CreateUserBasicRequest	true	"params"
//	@Success	200		{object}	v1.Response
//	@Router		/user_auto [post]
func (h *UserHandler) CreateUserBasic(ctx *gin.Context) {
	var req v1.CreateUserBasicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	// 创建用户基本信息
	if _, err := h.userService.CreateUserBasic(req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// LoginByOpenId godoc
//
//	@Summary	使用openid登录
//	@Schemes
//	@Description
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Param		openid	query string	true	"openid"
//	@Success	200		{object}	v1.LoginResponse
//	@Router		/login_openid [get]
func (h *UserHandler) LoginByOpenId(ctx *gin.Context) {
	openId := ctx.Query("openid")
	data, err := h.userService.LoginByOpenId(ctx, openId)

	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	v1.HandleSuccess(ctx, data)
}

// GetOpenId godoc
//
//	@Summary	获取openid
//	@Schemes
//	@Description	后台微信小程序获取openid
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Param		code	query string	true	"js_code"
//	@Success	200		{object}	wx.GetOpenIdByCodeResponse
//	@Router		/openid [get]
func (h *UserHandler) GetOpenId(ctx *gin.Context) {
	code := ctx.Query("code")

	data := wx.GetOpenIdByCode(code)
	v1.HandleSuccess(ctx, data)
}

// UpdateUserStudentCode godoc
//
//	@Summary	更新用户的学生代码
//	@Schemes
//	@Description	根据提供的学生代码和用户ID更新用户信息
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.UpdateUserStudentCode	true	"params"
//	@Security	Bearer
//	@Success	200		{object}	v1.Response
//	@Router		/user/studentcode [put]
func (h *UserHandler) UpdateUserStudentCode(ctx *gin.Context) {
	var req v1.UpdateUserStudentCode

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err := h.userService.UpdateUserStudentCode(ctx, req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// Logout godoc
//
//	@Summary	用户登出
//	@Schemes
//	@Description	根据提供的用户ID登出用户
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Param		userId	query string	true	"userId"
//	@Security	Bearer
//	@Success	200		{object}	v1.Response
//	@Router		/user/logout [get]
func (h *UserHandler) Logout(ctx *gin.Context) {
	userId := ctx.Query("userId")

	err := h.userService.Logout(ctx, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, "logout success")
}
