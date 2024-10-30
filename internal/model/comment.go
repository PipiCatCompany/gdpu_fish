package model

import "gorm.io/gorm"

// 上传商品评论表
type Comment struct {
	gorm.Model
	PostId  uint   `gorm:"type:bigint" json:"post_id"` // 评论父实例id
	UserId  string `gorm:"type:bigint" json:"user_id"`
	Content string `gorm:"type:string" json:"content"`
}

func (m *Comment) TableName() string {
	return "comment"
}
