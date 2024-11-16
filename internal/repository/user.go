package repository

import (
	"context"
	"errors"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByOpenId(ctx context.Context, openid string) (*model.User, error)
	CreateUserBasic(u model.User) (*model.User, error)
	GetUserCommentProfile(userId string) (v1.UserCommentProfile, error)
	SetUserOnline(ctx context.Context, userId string) error
	SetUserOffline(ctx context.Context, userId string) error
}

func NewUserRepository(
	r *Repository,
) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByOpenId(ctx context.Context, openid string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("open_id = ?", openid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// openid第一次登陆 创建Basic_User
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUserBasic(u model.User) (*model.User, error) {
	result := r.db.Create(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

// 通过userId 获取用户信息
func (r *userRepository) GetUserCommentProfile(userId string) (v1.UserCommentProfile, error) {
	var user model.User
	if err := r.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return v1.UserCommentProfile{}, err
		}
		return v1.UserCommentProfile{}, err
	}
	return v1.UserCommentProfile{
		Username: user.Nickname,
		Avatar:   user.Avatar,
	}, nil
}

// 同步用户登陆状态 -> Redis
func (r *userRepository) SetUserOnline(ctx context.Context, userId string) error {
	result := r.rdb.Set(ctx, userId, "online", 0).Err()
	return result
}

func (r *userRepository) SetUserOffline(ctx context.Context, userId string) error {
	result := r.rdb.Del(ctx, userId).Err()
	return result
}
