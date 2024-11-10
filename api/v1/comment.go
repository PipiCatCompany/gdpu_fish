package v1

type CreateCommentRequest struct {
	PostId  uint   `json:"post_id"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
}

type CommentResponse struct {
	CreateTime string `json:"create_time"`
	Username   string `json:"username"`
	Avatar     string `json:"avatar"`
	Content    string `json:"content"`
}
