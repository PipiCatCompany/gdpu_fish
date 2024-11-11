package model

import "gorm.io/gorm"

type Chatroom struct {
	gorm.Model
	ChatRoomId string `gorm:"unique;not null"` // postId-sellerId-buyerId
}

func (m *Chatroom) TableName() string {
	return "chatroom"
}
