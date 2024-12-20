package v1

import "go-xianyu/internal/model"

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" example:"alan"`
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
}
type GetProfileResponseData struct {
	UserId   string `json:"userId"`
	Nickname string `json:"nickname" example:"alan"`
}
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}

type CreateUserBasicRequest struct {
	// 默认是openid
	Username string `json:"username"`
	OpenId   string `json:"openid"`
}

type UserCommentProfile struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type LoginByOpenidResponse struct {
	User        model.User `json:"user"`
	AccessToken string     `json:"accessToken"`
}

type UpdateUserStudentCode struct {
	UserId      string `json:"userId"`
	StudentCode string `json:"studentCode"`
}
