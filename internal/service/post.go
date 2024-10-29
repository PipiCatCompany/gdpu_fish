package service

import (
	"context"
	"go-xianyu/internal/model"
	"go-xianyu/internal/repository"
)

type PostService interface {
	GetPost(ctx context.Context, id int64) (*model.Post, error)
	CreatePost(post *model.Post) error
}

func NewPostService(
	service *Service,
	postRepository repository.PostRepository,
) PostService {
	return &postService{
		Service:        service,
		postRepository: postRepository,
	}
}

type postService struct {
	*Service
	postRepository repository.PostRepository
}

func (s *postService) GetPost(ctx context.Context, id int64) (*model.Post, error) {
	return s.postRepository.GetPost(ctx, id)
}

func (s *postService) CreatePost(post *model.Post) error {
	return s.postRepository.Create(post)
}
