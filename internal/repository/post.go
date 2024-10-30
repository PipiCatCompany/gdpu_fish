package repository

import (
	"context"
	"go-xianyu/internal/model"
)

type PostRepository interface {
	GetPost(ctx context.Context, id int64) (*model.Post, error)
	Create(post *model.Post) error
	PostPagination(pageNum int, pageSize int) ([]model.Post, error)
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

func (r *postRepository) PostPagination(pageNum int, pageSize int) ([]model.Post, error) {
	offset := (pageNum - 1) * pageSize
	posts := make([]model.Post, pageSize)
	result := r.db.Offset(offset).Limit(pageSize).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}
