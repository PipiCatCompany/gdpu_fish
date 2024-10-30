package v1

type CreateCommentRequest struct {
	PostId  uint   `json:"post_id"`
	UserId  uint   `json:"user_id"`
	Content string `json:"content"`
}

type CommentResponse struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Content  string `json:"content"`
}
