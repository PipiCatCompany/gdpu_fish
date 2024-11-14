package v1

import "time"

type PostPaginationResponse struct {
	PostId     uint      `json:"post_id"`
	Title      string    `json:"title"`
	Info       string    `json:"info"`
	Username   string    `json:"username"`
	SellerId   string    `join:"seller_id"`
	UserAvatar string    `json:"user_avatar"`
	Price      float64   `json:"price"`
	Img        []string  `json:"img"`
	Category   string    `json:"category"`
	CreateTime time.Time `json:"createtime"`
}

type CreatePostRequest struct {
	Title    string  `json:"title"`
	Info     string  `json:"info"`
	UserId   string  `json:"userId"`
	Price    float64 `json:"price"`
	Img      string  `json:"img"`
	Category string  `json:"category"`
}
