package repository

import (
    "context"
	"go-xianyu/internal/model"
)

type CommentRepository interface {
	GetComment(ctx context.Context, id int64) (*model.Comment, error)
}

func NewCommentRepository(
	repository *Repository,
) CommentRepository {
	return &commentRepository{
		Repository: repository,
	}
}

type commentRepository struct {
	*Repository
}

func (r *commentRepository) GetComment(ctx context.Context, id int64) (*model.Comment, error) {
	var comment model.Comment

	return &comment, nil
}
