package model

type Message struct {
	ID        uint   `gorm:"primarykey"`
	CreatedAt string `gorm:"type:string" json:"create_time"`
	// ChatroomId = postId-sellerId-buyerId -> 查找会导致索引失效
	PostId   string `gorm:"type:string;" json:"post_id"`
	SellerId string `gorm:"type:string;index" json:"seller_id"`
	BuyerId  string `gorm:"type:string;index" json:"buyer_id"`
	Content  string `gorm:"type:string" json:"content"`
	Read     bool   `gorm:"type:bool" json:"read"`
}

func (m *Message) TableName() string {
	return "message"
}
