package service

import (
	"context"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
	"go-xianyu/internal/repository"
)

type MessageService interface {
	GetMessage(ctx context.Context, id int64) (*model.Message, error)
	CreateMessage(message v1.CreateMessageRequest) error
}

func NewMessageService(
	service *Service,
	messageRepository repository.MessageRepository,
) MessageService {
	return &messageService{
		Service:           service,
		messageRepository: messageRepository,
	}
}

type messageService struct {
	*Service
	messageRepository repository.MessageRepository
}

func (s *messageService) GetMessage(ctx context.Context, id int64) (*model.Message, error) {
	return s.messageRepository.GetMessage(ctx, id)
}

func (s *messageService) CreateMessage(message v1.CreateMessageRequest) error {
	return s.messageRepository.CreateMessage(message)
}
