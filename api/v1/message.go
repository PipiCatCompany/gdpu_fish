package v1

type CreateMessageRequest struct {
	CreatedAt string `json:"create_time"`
	PostId    string `json:"post_id"`
	SellerId  string `json:"seller_id"`
	BuyerId   string `json:"buyer_id"`
	Content   string `json:"content"`
	Read      bool   `json:"read"`
}
