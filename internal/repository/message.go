package repository

import (
	"context"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
)

type MessageRepository interface {
	GetMessage(ctx context.Context, id int64) (*model.Message, error)
	CreateMessage(message v1.CreateMessageRequest) error
}

func NewMessageRepository(
	repository *Repository,
) MessageRepository {
	return &messageRepository{
		Repository: repository,
	}
}

type messageRepository struct {
	*Repository
}

func (r *messageRepository) GetMessage(ctx context.Context, id int64) (*model.Message, error) {
	var message model.Message

	return &message, nil
}

func (r *messageRepository) CreateMessage(message v1.CreateMessageRequest) error {
	return r.Repository.db.Create(message).Error
}
