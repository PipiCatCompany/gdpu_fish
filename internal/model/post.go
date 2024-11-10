package model

import "gorm.io/gorm"

// 上传商品表
type Post struct {
	gorm.Model
	Title    string  `gorm:"type:string" json:"title"`    // 商品标题
	Info     string  `gorm:"type:string" json:"info"`     // 商品说明
	UserId   string  `gorm:"type:bigint" json:"userId"`   // 发布者UserId
	Price    float64 `gorm:"type:float" json:"price"`     // 商品价格
	Img      string  `gorm:"type:string" json:"img"`      // 商品图片URL - 多图用分号隔开
	Category string  `gorm:"type:string" json:"category"` // 类别：二手 / 兼职任务
}

func (m *Post) TableName() string {
	return "post"
}
