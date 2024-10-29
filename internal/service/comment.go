package service

import (
    "context"
	"go-xianyu/internal/model"
	"go-xianyu/internal/repository"
)

type CommentService interface {
	GetComment(ctx context.Context, id int64) (*model.Comment, error)
}
func NewCommentService(
    service *Service,
    commentRepository repository.CommentRepository,
) CommentService {
	return &commentService{
		Service:        service,
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
