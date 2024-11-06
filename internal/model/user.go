package model

import (
	"gorm.io/gorm"
)

// 用户表
type User struct {
	gorm.Model
	UserId      string `gorm:"unique;not null"`
	Nickname    string `gorm:"null"`
	Password    string `gorm:"null"`
	OpenId      string `gorm:"unique;not null"`
	Email       string `gorm:"null"`
	StudentCode string `gorm:"null"`
	Avatar      string `gorm:"null"` // 用户头像
}

func (u *User) TableName() string {
	return "users"
}
