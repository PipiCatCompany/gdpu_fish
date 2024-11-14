package v1

type CreateMessageRequest struct {
	CreatedAt string `json:"create_time"`
	PostId    uint   `json:"post_id"`
	SellerId  string `json:"seller_id"`
	BuyerId   string `json:"buyer_id"`
	Content   string `json:"content"`
	Read      bool   `json:"read"`
	MsgSender string `json:"msg_sender"` // 发消息的用户id
}

type MessageChanelResponse struct {
	StuffInfo interface{} `json:"stuff_info"`
	UserInfo  interface{} `json:"user_info"`
}
