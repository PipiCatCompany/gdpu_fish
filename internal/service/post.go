package service

import (
	"context"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
	"go-xianyu/internal/repository"
)

type PostService interface {
	GetPost(ctx context.Context, id int64) (*model.Post, error)
	CreatePost(post *model.Post) error
	GetPostListByPage(pageNum int, pageSize int) ([]v1.PostPaginationResponse, error)
}

func NewPostService(
	service *Service,
	postRepository repository.PostRepository,
	userRepository repository.UserRepository,
) PostService {
	return &postService{
		Service:        service,
		postRepository: postRepository,
		userRepository: userRepository,
	}
}

type postService struct {
	*Service
	postRepository repository.PostRepository
	userRepository repository.UserRepository
}

func (s *postService) GetPost(ctx context.Context, id int64) (*model.Post, error) {
	return s.postRepository.GetPost(ctx, id)
}

func (s *postService) CreatePost(post *model.Post) error {
	return s.postRepository.Create(post)
}

func (s *postService) GetPostListByPage(pageNum int, pageSize int) ([]v1.PostPaginationResponse, error) {
	posts, err := s.postRepository.PostPagination(pageNum, pageSize)
	data := make([]v1.PostPaginationResponse, 0)

	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		user, err := s.userRepository.GetUserCommentProfile(post.UserId)
		if err != nil {
			return nil, err
		}

		data = append(data, v1.PostPaginationResponse{
			PostId:     post.ID,
			Title:      post.Title,
			Info:       post.Info,
			Price:      post.Price,
			Img:        post.Img,
			Category:   post.Category,
			CreateTime: post.CreatedAt,
			Username:   user.Username,
			UserAvatar: user.Avatar,
		})
	}

	return data, nil
}
