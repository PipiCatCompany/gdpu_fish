package service

import (
	"context"
	v1 "go-xianyu/api/v1"
	"go-xianyu/internal/model"
	"go-xianyu/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
	CreateUserBasic(req v1.CreateUserBasicRequest) (*model.User, error)
	LoginByOpenId(ctx context.Context, openid string) (v1.LoginByOpenidResponse, error)
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check username
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if err == nil && user != nil {
		return v1.ErrEmailAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userId, err := s.sid.GenString()
	if err != nil {
		return err
	}
	user = &model.User{
		UserId:   userId,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a user
		if err = s.userRepo.Create(ctx, user); err != nil {
			return err
		}
		// TODO: other repo
		return nil
	})
	return err
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return "", v1.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	// 用userId生成Jwt令牌
	token, err := s.jwt.GenToken(user.UserId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) LoginByOpenId(ctx context.Context, openid string) (v1.LoginByOpenidResponse, error) {
	user, err := s.userRepo.GetByOpenId(ctx, openid)
	if err != nil || user == nil {
		return v1.LoginByOpenidResponse{}, v1.ErrUnauthorized
	}
	// openid 只会返回对应一个用户的token，所以不需要额外验证
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	// if err != nil {
	// 	return "", err
	// }

	// 用userId生成Jwt令牌
	token, err := s.jwt.GenToken(user.UserId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return v1.LoginByOpenidResponse{}, err
	}

	// 登陆返回用户信息给前端持久化
	return v1.LoginByOpenidResponse{User: *user, AccessToken: token}, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserId:   user.UserId,
		Nickname: user.Nickname,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) CreateUserBasic(req v1.CreateUserBasicRequest) (*model.User, error) {
	userId, err := s.sid.GenString()
	if err != nil {
		return nil, err
	}

	return s.userRepo.CreateUserBasic(model.User{
		UserId:      userId, // 用户雪花id
		Nickname:    req.Username,
		Password:    "",
		OpenId:      req.OpenId,
		Email:       "",
		StudentCode: "",
	})
}
