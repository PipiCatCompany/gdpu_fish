package repository

import (
	"context"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
)

type CommentRepository interface {
	GetComment(ctx context.Context, id int64) (*model.Comment, error)
	CreateComment(comment *model.Comment) error
	GetCommentList(postId uint) ([]v1.CommentResponse, error)
}

func NewCommentRepository(
	repository *Repository,
	userRepository UserRepository,
) CommentRepository {
	return &commentRepository{
		repository:     repository,
		userRepository: userRepository,
	}
}

type commentRepository struct {
	repository     *Repository
	userRepository UserRepository // 依赖注入
}

func (r *commentRepository) GetComment(ctx context.Context, id int64) (*model.Comment, error) {
	var comment model.Comment

	return &comment, nil
}

func (r *commentRepository) CreateComment(comment *model.Comment) error {
	return r.repository.db.Create(comment).Error
}

func (r *commentRepository) GetCommentList(postId uint) ([]v1.CommentResponse, error) {
	// 获取postId下的评论
	var comments []model.Comment

	result := r.repository.db.Where("post_id = ?", postId).Find(&comments)

	if result.Error != nil {
		return []v1.CommentResponse{}, result.Error
	}

	commentList := make([]v1.CommentResponse, 0)
	for _, comment := range comments {
		// 序列化用户信息
		userId := comment.UserId
		userProfile, _ := r.userRepository.GetUserCommentProfile(userId)

		commentList = append(commentList, v1.CommentResponse{
			CreateTime: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			Username:   userProfile.Username,
			Avatar:     userProfile.Avatar,
			Content:    comment.Content,
		})
	}

	return commentList, nil
}
