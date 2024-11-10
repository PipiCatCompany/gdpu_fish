package model

import "gorm.io/gorm"

type Chatroom struct {
	gorm.Model
	ChatRoomId string `gorm:"unique;not null"` //
}

func (m *Chatroom) TableName() string {
	return "chatroom"
}
