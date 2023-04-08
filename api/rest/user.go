package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iokr/bbs/api/request"
	"github.com/iokr/bbs/internal/biz"
	"github.com/iokr/bbs/internal/response"
)

type UserRest struct {
	userBiz *biz.UserBiz
}

func NewUserRest(userBiz *biz.UserBiz) *UserRest {
	return &UserRest{userBiz: userBiz}
}

func (r *UserRest) Register(ctx *gin.Context) {
	var param request.RegisterRequest

	err := ctx.MustBindWith(&param, binding.JSON)
	if err != nil {
		response.ServerJson(ctx.Writer, nil, err)
		return
	}
	response.ServerJson(ctx.Writer, nil, r.userBiz.Register(ctx, &param))
}

func (r *UserRest) Login(ctx *gin.Context) {
	var param request.LoginRequest

	err := ctx.MustBindWith(&param, binding.JSON)
	if err != nil {
		response.ServerJson(ctx.Writer, nil, err)
		return
	}

	result, err := r.userBiz.Login(ctx, &param)
	if err != nil {
		response.ServerJson(ctx.Writer, nil, err)
		return
	}
	response.ServerJson(ctx.Writer, result, err)
}
