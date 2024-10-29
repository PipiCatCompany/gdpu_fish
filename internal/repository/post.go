package repository

import (
	"context"
	"go-xianyu/internal/model"
)

type PostRepository interface {
	GetPost(ctx context.Context, id int64) (*model.Post, error)
	Create(post *model.Post) error
}

func NewPostRepository(
	repository *Repository,
) PostRepository {
	return &postRepository{
		Repository: repository,
	}
}

type postRepository struct {
	*Repository
}

func (r *postRepository) GetPost(ctx context.Context, id int64) (*model.Post, error) {
	var post model.Post

	return &post, nil
}

func (r *postRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}
