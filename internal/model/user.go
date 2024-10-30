package model

import (
	"gorm.io/gorm"
)

// 用户表
type User struct {
	gorm.Model
	UserId      string `gorm:"unique;not null"`
	Nickname    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	OpenId      string `gorm:"unique;null"`
	Email       string `gorm:"unique;not null"`
	StudentCode string `gorm:"unique;null"`
	Avatar      string `gorm:"null"` // 用户头像
}

func (u *User) TableName() string {
	return "users"
}
