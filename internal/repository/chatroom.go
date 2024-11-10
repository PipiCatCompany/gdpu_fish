package repository

import (
    "context"
	"go-xianyu/internal/model"
)

type ChatroomRepository interface {
	GetChatroom(ctx context.Context, id int64) (*model.Chatroom, error)
}

func NewChatroomRepository(
	repository *Repository,
) ChatroomRepository {
	return &chatroomRepository{
		Repository: repository,
	}
}

type chatroomRepository struct {
	*Repository
}

func (r *chatroomRepository) GetChatroom(ctx context.Context, id int64) (*model.Chatroom, error) {
	var chatroom model.Chatroom

	return &chatroom, nil
}
