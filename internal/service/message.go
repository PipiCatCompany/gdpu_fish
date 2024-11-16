package service

import (
	"context"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
	"go-xianyu/internal/repository"
	"go-xianyu/pkg/websocket"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	conf *viper.Viper,
) MessageService {
	return &messageService{
		Service:           service,
		messageRepository: messageRepository,
		postRepository:    postRepository,
		userRepository:    userRepository,
		conf:              conf,
	}
}

type messageService struct {
	*Service
	messageRepository repository.MessageRepository
	postRepository    repository.PostRepository
	userRepository    repository.UserRepository
	conf              *viper.Viper
}

func (s *messageService) GetMessage(ctx context.Context, id int64) (*model.Message, error) {
	return s.messageRepository.GetMessage(ctx, id)
}

func (s *messageService) CreateMessage(message v1.CreateMessageRequest) error {
	// 增加异步操作
	// go func(conf *viper.Viper, message v1.CreateMessageRequest) {
	// 	// 转发给Cpp长连接服务器
	// 	errCppServer := websocket.SyncMessageToCpp(conf, message)
	// 	if errCppServer != nil {
	// 		fmt.Printf("%v", errCppServer)
	// 	}
	// }(s.conf, message)

	websocket.SyncMessageToCpp(s.conf, message)

	return s.messageRepository.CreateMessage(message)
}

func (s *messageService) GetMessageListByPage(pageSize int, pageNum int, chartroomId string) ([]model.Message, error) {
	return s.messageRepository.GetMessageListByPage(pageSize, pageNum, chartroomId)
}

func (s *messageService) GetMessageChanelInfo(ctx *gin.Context, chatroomId string) (v1.MessageChanelResponse, error) {
	list := strings.Split(chatroomId, "-")
	postId, userId1, userId2 := list[0], list[1], list[2]
	postIdI, _ := strconv.ParseInt(postId, 10, 64)

	post, err := s.postRepository.GetPostById(uint(postIdI))
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
			// userId 为Key的Map
			userId1: map[string]interface{}{"username": user1.Username, "avatar": user1.Avatar},
			userId2: map[string]interface{}{"username": user2.Username, "avatar": user2.Avatar},
		},
	}, nil
}
