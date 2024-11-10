package service

import (
    "context"
	"go-xianyu/internal/model"
	"go-xianyu/internal/repository"
)

type ChatroomService interface {
	GetChatroom(ctx context.Context, id int64) (*model.Chatroom, error)
}
func NewChatroomService(
    service *Service,
    chatroomRepository repository.ChatroomRepository,
) ChatroomService {
	return &chatroomService{
		Service:        service,
		chatroomRepository: chatroomRepository,
	}
}

type chatroomService struct {
	*Service
	chatroomRepository repository.ChatroomRepository
}

func (s *chatroomService) GetChatroom(ctx context.Context, id int64) (*model.Chatroom, error) {
	return s.chatroomRepository.GetChatroom(ctx, id)
}
