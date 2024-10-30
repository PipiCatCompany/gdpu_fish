package v1

import "time"

type PostPaginationResponse struct {
	Info string `json:"info"`
	// UserId     string    `json:"userId"`
	Username   string    `json:"username"`
	UserAvatar string    `json:"user_avatar"`
	Price      float64   `json:"price"`
	Img        string    `json:"img"`
	Category   string    `json:"category"`
	CreateTime time.Time `json:"createtime"`
}
