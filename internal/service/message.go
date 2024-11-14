package service

import (
	"context"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
	"go-xianyu/internal/repository"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MessageService interface {
	GetMessage(ctx context.Context, id int64) (*model.Message, error)
	CreateMessage(message v1.CreateMessageRequest) error
	GetMessageListByPage(pageSize int, pageNum int, chartroomId string) ([]model.Message, error)
	GetMessageChanelInfo(ctx *gin.Context, chatroomId string) (v1.MessageChanelResponse, error)
}

func NewMessageService(
	service *Service,
	messageRepository repository.MessageRepository,
	postRepository repository.PostRepository,
	userRepository repository.UserRepository,
) MessageService {
	return &messageService{
		Service:           service,
		messageRepository: messageRepository,
		postRepository:    postRepository,
		userRepository:    userRepository,
	}
}

type messageService struct {
	*Service
	messageRepository repository.MessageRepository
	postRepository    repository.PostRepository
	userRepository    repository.UserRepository
}

func (s *messageService) GetMessage(ctx context.Context, id int64) (*model.Message, error) {
	return s.messageRepository.GetMessage(ctx, id)
}

func (s *messageService) CreateMessage(message v1.CreateMessageRequest) error {
	return s.messageRepository.CreateMessage(message)
}

func (s *messageService) GetMessageListByPage(pageSize int, pageNum int, chartroomId string) ([]model.Message, error) {
	return s.messageRepository.GetMessageListByPage(pageSize, pageNum, chartroomId)
}

func (s *messageService) GetMessageChanelInfo(ctx *gin.Context, chatroomId string) (v1.MessageChanelResponse, error) {
	list := strings.Split(chatroomId, "-")
	postId, userId1, userId2 := list[0], list[1], list[2]
	postIdI, _ := strconv.ParseInt(postId, 10, 64)

	post, err := s.postRepository.GetPostById(postIdI)
	if err != nil {
		return v1.MessageChanelResponse{}, err
	}

	user1, err := s.userRepository.GetUserCommentProfile(userId1)
	if err != nil {
		return v1.MessageChanelResponse{}, err
	}

	user2, err := s.userRepository.GetUserCommentProfile(userId2)
	if err != nil {
		return v1.MessageChanelResponse{}, err
	}

	// TODO 可以优化
	return v1.MessageChanelResponse{
		StuffInfo: post,
		UserInfo: map[string]interface{}{
			"seller": map[string]interface{}{"userId": userId1, "avatar": user1.Avatar},
			"buyer":  map[string]interface{}{"userId": userId2, "avatar": user2.Avatar},
		},
	}, nil
}
