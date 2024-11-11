package v1

type CreateMessageRequest struct {
	CreatedAt  string `gorm:"type:string" json:"create_time"`
	ChatroomId string `gorm:"type:string" json:"chatroomId"`
	Content    string `gorm:"type:string" json:"content"`
	Read       bool   `json:"read"`
}
