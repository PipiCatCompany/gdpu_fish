package model

type Message struct {
	ID         uint   `gorm:"primarykey"`
	CreatedAt  string `gorm:"type:string" json:"create_time"`
	ChatroomId string `gorm:"type:string" json:"chatroomId"`
	Content    string `gorm:"type:string" json:"content"`
	Read       bool   `gorm:"type:bool" json:"read"`
	// 极简的Message，其他字段的序列化给到imService调Base服务实现
}

func (m *Message) TableName() string {
	return "message"
}
