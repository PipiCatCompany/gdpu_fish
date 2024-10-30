package service

import (
	"context"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
	"go-xianyu/internal/repository"
)

type CommentService interface {
	GetComment(ctx context.Context, id int64) (*model.Comment, error)
	GetCommentList(postId uint) ([]v1.CommentResponse, error)
}

func NewCommentService(
	service *Service,
	commentRepository repository.CommentRepository,
) CommentService {
	return &commentService{
		Service:           service,
		commentRepository: commentRepository,
	}
}

type commentService struct {
	*Service
	commentRepository repository.CommentRepository
}

func (s *commentService) GetComment(ctx context.Context, id int64) (*model.Comment, error) {
	return s.commentRepository.GetComment(ctx, id)
}

func (s *commentService) CreateComment(req v1.CreateCommentRequest) error {
	return s.commentRepository.CreateComment(
		&model.Comment{
			PostId:  req.PostId,
			UserId:  req.UserId,
			Content: req.Content,
		},
	)
}

func (s *commentService) GetCommentList(postId uint) ([]v1.CommentResponse, error) {
	return s.commentRepository.GetCommentList(postId)
}
