package model

type Message struct {
	ID         uint   `gorm:"primarykey"`
	CreatedAt  string `gorm:"type:string" json:"create_time"`
	PostId     uint   `gorm:"type:string;" json:"post_id"`
	SellerId   string `gorm:"type:string;index" json:"seller_id"`
	BuyerId    string `gorm:"type:string;index" json:"buyer_id"`
	ChatroomId string `gorm:"type:string;index" json:"chatroom_id"`
	Content    string `gorm:"type:string" json:"content"`
	Read       bool   `gorm:"type:bool" json:"read"`
	MsgSender  string `gorm:"type:string" json:"msg_sender"` // 发消息的用户id
}

// ChatroomId = postId-sellerId-buyerId -> 查找会导致索引失效
// 所以拆分成三个字段
// 33-DPpeAHgUF7-DJhYpqApGk

func (m *Message) TableName() string {
	return "message"
}
