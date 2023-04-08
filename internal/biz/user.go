package biz

import (
	"context"

	"github.com/iokr/bbs/api/request"
	"github.com/iokr/bbs/core/utils"
	"github.com/iokr/bbs/internal/data"
	"github.com/iokr/bbs/internal/errors"
	"github.com/iokr/bbs/internal/middleware"
	"github.com/samber/lo"
)

type UserBiz struct {
	userRepo *data.UserRepo
}

func NewUserBiz(userRepo *data.UserRepo) *UserBiz {
	return &UserBiz{
		userRepo: userRepo,
	}
}

func (l *UserBiz) Register(ctx context.Context, userReq *request.RegisterRequest) error {
	if userReq.UserName == "" {
		return errors.UserNameIsEmpty
	}
	if userReq.Email == "" {
		return errors.EmailIsEmpty
	}
	if userReq.Password == "" {
		return errors.PasswordIsEmpty
	}

	has, err := l.userRepo.ExistByEmail(ctx, userReq.Email)
	if err != nil {
		return err
	}
	if has {
		return errors.EmailIsExist
	}

	return l.userRepo.Create(ctx, &data.User{
		UserName: userReq.UserName,
		Email:    userReq.Email,
		Password: utils.MD5(lo.Reverse([]byte(userReq.Password))),
	})
}

func (l *UserBiz) Login(ctx context.Context, userReq *request.LoginRequest) (*request.LoginReply, error) {
	if userReq.Email == "" {
		return nil, errors.EmailIsEmpty
	}

	user, err := l.userRepo.GetByEmail(ctx, userReq.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.UserNotFound
	}

	if user.IsLocked == 1 {
		return nil, errors.EmailIsLocked
	}

	if utils.MD5(lo.Reverse([]byte(userReq.Password))) != user.Password {
		return nil, errors.PasswordIsError
	}

	claims := make(map[string]interface{})
	claims["userId"] = user.ID
	claims["email"] = user.Email
	token, err := middleware.GenerateToken(claims)
	if err != nil {
		return nil, err
	}
	return &request.LoginReply{
		Token: token,
	}, nil
}
